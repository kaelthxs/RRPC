<template>
    <div class="main">
        <div class="container">
            {{ roleOfAuthUser }}
            <button v-if="roleOfAuthUser == 'admin'" @click="funcCreateNewProductForProductComp">Добавить</button>
            <button @click="$emit('back')">Назад</button>
            
            <div class="main__header">
                Товары
            </div>
            <div class="main__underHeader">
                <div class="main__sidebar">
                    <div class="sidebar__title">
                        Фильтры
                    </div>
                    <div class="sidebar__pricepicker">
                        <div class="pricepicker__title">
                            Цена
                        </div>
                        <div class="price-range">
                            <input type="range" v-model="minPrice" :min="minPossiblePrice" :max="maxPossiblePrice" step="500" class="price-range__range1">
                            <input type="range" v-model="maxPrice" :min="minPossiblePrice" :max="maxPossiblePrice" step="500" class="price-range__range2">
                        </div>
                        <div class="price-range__texts">
                            <input type="text" v-model="minPrice" class="price-range__text">
                            <input type="text" v-model="maxPrice" class="price-range__text">
                        </div>
                    </div>
                    <div class="sidebar__creatorpicker">
                        <div class="creatopicker__title">
                            Производитель
                        </div>
                        <div class="creatopicker__inputs">
                            <label class="list-check__label" v-for="creator in uniqueCreators" :key="creator">
                            <input type="checkbox" :value="creator" v-model="selectedCreators" class="creatopicker__inputs--example">
                            <span class="label-text"> {{ creator }} </span>
                            </label>
                        </div>

                    </div>
                </div>
                <div class="product__item--new" v-if="isCreateNew">
                    <input type="text" v-model="createData.Creator" placeholder="creator">
                    <input type="text" v-model="createData.Name" placeholder="name">
                    <input type="number" v-model="createData.Price" placeholder="price">
                    <input type="text" v-model="createData.Description" placeholder="desc">
                    <input type="number" v-model="createData.Stock" placeholder="stock">

                    <button @click="sendNewProduct">Сохранить</button>
                </div>
                <div class="main__main" v-if="filteredProducts.length">
                    <div v-for="product in filteredProducts" :key="product.id" class="product__item">
                        <input v-model="getEditableProduct(product.id).Creator" type="text" v-if="isRewriteSelectedProductsForCreateNewProduct">
                        <input v-model="getEditableProduct(product.id).Name" type="text" v-if="isRewriteSelectedProductsForCreateNewProduct">

                        <div class="item__title">
                            <span v-if="!isRewriteSelectedProductsForCreateNewProduct">{{ product.Creator + " " + product.Name }}</span>
                            <button v-if="roleOfAuthUser == 'admin'" @click="prepareEdit(product)">Изменить</button>
                        </div>
                        <input v-model="getEditableProduct(product.id).Price" type="number" v-if="isRewriteSelectedProductsForCreateNewProduct">
                        <div class="item__price" v-if="!isRewriteSelectedProductsForCreateNewProduct">
                            {{ product.Price }}₽
                        </div>
                        <button @click="addToCart(product.id)">Добавить в корзину</button>
                        <button v-if="editProduct[product.id]" @click="saveProduct(product.id, product.CategoryID)">Сохранить</button>
                        <button v-if="roleOfAuthUser == 'admin'" @click="deleteProduct(product.id)">Удалить</button>
                    </div>


            </div>
                <div v-else>
                    <p>Товары не найдены.</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script>

import axios from 'axios';
export default {
    props: {
        categoryId: {
            type: Number,
            required: true,
        }
    },
    computed: {

                // Определение минимальной и максимальной возможной цены среди всех товаров
        minPossiblePrice() {
            return this.products.length ? Math.min(...this.products.map(p => p.Price)) : 0;
        },
        maxPossiblePrice() {
            return this.products.length ? Math.max(...this.products.map(p => p.Price)) : 100000;
        },

        // Создаем список уникальных производителей
        uniqueCreators() {
            return [...new Set(this.products.map(product => product.Creator))].filter(Boolean);
        },

        filteredProducts() {
            return this.products.filter(product => {
                const matchesCreator = this.selectedCreators.length === 0 || this.selectedCreators.includes(product.Creator);
                const matchesPrice = product.Price >= this.minPrice && product.Price <= this.maxPrice;
                return matchesCreator && matchesPrice;
            });
        }
    },
    name: 'ProductsComp',
    data() {
        return {
            products: [],
            allProducts: [],
            cart: [],
            selectedCreators: [], // Выбранные производители
            minPrice: 0, // Нижняя граница цены
            maxPrice: 450000, // Верхняя граница цены
            roleOfAuthUser: '',
            createData:{
                Name: '',
                Creator: '',
                Price: '',
                Description: '',
                Stock: '',
                CategoryID: this.categoryId
            },
            isCreateNew: false,
            isRewrite: false,
            isEdit: false,
            editProduct: {},
            isRewriteSelectedProductsForCreateNewProduct: false
        }
    },

    watch: {
    categoryId(newVal) {
      if (newVal) {
        this.getProductsByCategory();
      }
    }
  },
  methods: {
    prepareEdit(product) {
        this.isRewriteSelectedProductsForCreateNewProduct = !this.isRewriteSelectedProductsForCreateNewProduct
        if (!this.editProduct[product.id]) {
            this.editProduct[product.id] = { ...product };
        }
    },

    getEditableProduct(id) {
        return this.editProduct[id] || { Name: "" };
    },

    async saveProduct(productId) {
        console.log(this.products);
        
        const authToken = localStorage.getItem('authToken');
        if (!authToken) {
            console.warn("Нет токена авторизации.");
            return;
        }

        const updatedData = {};
        for (const key in this.editProduct[productId]) {
            if (this.editProduct[productId][key] !== '') {
                updatedData[key] = this.editProduct[productId][key];
            }
        }

        try {
            updatedData['CategoryID'] = this.categoryId
            updatedData['Created_at'] = null
            console.log(updatedData);
            
            const response = await axios.put(
                `http://localhost:8080/api/product/${productId}`,
                JSON.stringify(updatedData),
                {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${authToken}`
                    }
                }
            );
            console.log('Продукт обновлён:', response.data);
            delete this.editProduct[productId];
        } catch (error) {
            console.error('Ошибка при обновлении продукта:', error);
        }
    },
    async deleteProduct(productId) {
        try {
            const authToken = localStorage.getItem('authToken');
            if (!authToken) {
                console.warn("Нет токена авторизации.");
                return;
            }

            await axios.delete(`http://localhost:8080/api/product/${productId}`, {
                headers: {
                    'Authorization': `Bearer ${authToken}`
                }
            });

            // Убираем товар из локального массива после удаления
            this.products = this.products.filter(product => product.id !== productId);

            
            console.log(`Товар с ID ${productId} удален.`);
        } catch (error) {
            console.error('Ошибка при удалении продукта:', error);
        }
    },
    addToCart (id) {
        console.log(id);
        console.log(this.allProducts);
        for (const el of this.allProducts) {
            if (el['ID'] == id) {
                this.cart.push(el)
            }
        }
        console.log('ВЫВОД КОРЗИНЫ НЕ ИЗ LS');
        
        console.log(this.cart);
        console.log(JSON.stringify(this.cart));
        
        console.log(`----------`);

        localStorage.setItem('cart', JSON.stringify(this.cart))

        console.log(`------------`);
    },
    async getProductsByCategory() {
        console.log("Получен categoryId:", this.categoryId);

        if (!this.categoryId) {
            console.warn("Ошибка: categoryId пустой!");
            return;
        }

        try {
            const authToken = localStorage.getItem('authToken');
            if (!authToken) {
                console.warn("Нет токена авторизации.");
                return;
            }

            const response = await axios.get(`http://localhost:8080/api/product/productsbycategory/${this.categoryId}`, {
                headers: {
                    'Authorization': `Bearer ${authToken}`
                }
            });

            console.log("Ответ сервера:", response.data);
            
            response.data.forEach(el => {
                this.allProducts.push(el)
            });

            // Доступ к массиву
            const productsArray = response.data; // Проверка на data в ответе
            console.log("Массив продуктов:", productsArray);

            if (Array.isArray(productsArray)) {
                this.products = productsArray.map(item => ({
                    id: item['ID'],
                    image: item.img,
                    Name: item.Name || "Без названия",
                    Creator: item.Creator || "Неизвестно",
                    Description: item.Description || "Описание отсутствует",
                    Price: item.Price || 0,
                    Stock: item.Stock || 0,
                    CategoryID: item.CategoryID,
                    Created_at: item.Created_at
                }));

                console.log("Товары успешно загружены:", this.products);
            } else {
                console.error("Ответ не содержит массива продуктов.");
            }

        } catch (error) {
            console.error("Ошибка при запросе:", error);
        }
    },
    getUserIdFromJWT() {
        const token = localStorage.getItem('authToken'); 
        if (!token) {
            console.warn("JWT токен не найден");
            return null;
        }

        try {
            const payloadBase64 = token.split('.')[1]; // Получаем payload (вторая часть JWT)
            const decodedPayload = JSON.parse(atob(payloadBase64)); // Декодируем из base64

            return decodedPayload.UserId;
            
        } catch (error) {
            console.error('Ошибка декодирования JWT:', error);
            return null;
        }
    },


    async getUserRoleById() {
        const userId = this.getUserIdFromJWT();        
        if (localStorage.getItem('authToken') != '') { 
            const response = await axios.get(`http://localhost:8080/api/user/${userId}`, {
            headers: {
            'Authorization': `Bearer ${localStorage.getItem('authToken')}`
            }   
            });
            // this.response = response.data;
            // console.log(this.response)
            this.roleOfAuthUser = response.data['Role'];
            console.log(this.roleOfAuthUser)
        }
    },
    funcCreateNewProductForProductComp () {
        this.isCreateNew = !this.isCreateNew
        this.isRewrite = !this.isRewrite
        console.log(this.isCreateNew);
        
    },

    reloadPage() {
      window.location.reload();
    },

    async sendNewProduct() {
        try {
            const authToken = localStorage.getItem('authToken');
            if (!authToken) {
                console.warn("Нет токена авторизации.");
                return;
            }

            const response = await axios.post(
                    'http://localhost:8080/api/product/', 
                JSON.stringify(this.createData),
                { 
                    headers: { 
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${authToken}`
                    } 
                }
            );

            console.log('Ответ сервера:', response.data);

            // Проверяем, есть ли в ответе новый токен (если сервер обновляет его)
            if (response.data.token) {
                localStorage.setItem('authToken', response.data.token);
                console.log('Новый токен сохранен:', response.data.token);
            }

            this.responseMessage = response.data.message;
            this.reloadPage()
            
        } catch (error) {
            console.error('Ошибка при отправке данных:', error);
        }
    }

    },

  mounted() {
        console.log("ProductsComp получил categoryId:", this.categoryId)
        this.getUserRoleById()
        this.getProductsByCategory()
    }   
}

</script>

<style scoped>

.price-range {
    margin-bottom: 20px;
    background-color: #272727;
    display: flex;
    position: relative;
}
.price-range__range1 {
    width: 200px;
    -webkit-appearance: none;
    -moz-appearance: none;
    appearance: none;
    position: absolute;
}

.price-range__range1::-webkit-slider-runnable-track {
    border-radius: 10px/100%;
    height: 3px;
    background-color: #5A5A5A;
}

.price-range__range1::-webkit-slider-thumb {
    background: #272727;
    border: 2px solid #5A5A5A;
    border-radius: 500%;
    width: 30px;
    height: 30px;
    -webkit-appearance: none;
    margin-bottom: 50px;
    transform: translateY(-45%);
}


.price-range__range2 {
    width: 200px;
    -webkit-appearance: none;
    -moz-appearance: none;
    appearance: none;
}

.price-range__range2::-webkit-slider-runnable-track {
    border-radius: 10px/100%;
    height: 3px;
    background-color: #5A5A5A;
}

.price-range__range2::-webkit-slider-thumb {
    background: #272727;
    border: 2px solid #5A5A5A;
    border-radius: 500%;
    width: 30px;
    height: 30px;
    -webkit-appearance: none;
    transform: translateY(-45%);
}
.price-range__text {
    margin-top: 20px;
    width: 90px;
    background-color: #323232;
    outline: none;
    border: none;
    height: 30px;
    color: white;
    padding-left: 10px;
    border-radius: 20px;
}

.price-range__texts {
    display: flex;
    width: 80%;
    justify-content: space-between;
}

.sidebar__title {
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    justify-items: center;
    color: white;
    border-bottom: 1px solid #5A5A5A;
    width: 100%;
}
.sidebar__pricepicker {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    color: white;
    width: 100%;
    border-bottom: 1px solid #5A5A5A;
    height: 30%;
    margin-bottom: 20px;

}

.pricepicker__title {
    margin-bottom: 20px;
}

.sidebar__creatorpicker {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    color: white;
    width: 80%;

}

.creatopicker__title {
    margin-bottom: 20px;

}
.creatopicker__inputs {
    display: flex;
    background-color: #222222;
    flex-direction: column;
    width: 80%;
    border-radius: 20px;
    padding: 10px;

}

.creatopicker__inputs--example::before {
    content: '';
    display: inline-block;
    border-radius: 1em;
    background-repeat: no-repeat;
    background-position: center center;
    background-color: #272727;
    height: 50px;
    width: 50px;

}

.creatopicker__inputs--example:checked::before {
    background-color: #5A5A5A;
    content: '';
}

.product__item {
    background-color: #2B2B2B;
    margin-top: 5px;
    width: 864px;
    height: 75px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 10px;
    color: white;
    transition: .3s;
}
.product__item--new {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    background-color: #2B2B2B;
    margin-top: 5px;
    width: 864px;
    height: 75px;
    flex-direction: column;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 10px;
    color: white;
    transition: .3s;
}

.product__item:hover {
    background-color: #323232;
}

.item__title {
}
.item__price {
}

.main__underHeader {
    width: 100%;
    display: flex;
}

.main {
    display: flex;
    align-items: center;
    justify-content: center;
    justify-items: center;
}
.container {   
    align-items: center;
    justify-content: center;
    width: 80%;
    justify-items: center;
}
.main__header {
    background-color: #323232;
    width: 100%;
    height: 50px;
    color: white;
    display: flex;
    justify-content: center;
    justify-items: center;
    align-items: center;
    border-radius: 30px 30px 0px 0px;
}
.main__sidebar {
    display: flex;
    width: 25%;
    background-color: #272727;
    flex-direction: column;
    align-items: center;
    height: 600px;
}
.main__main {
    display: flex;
    flex-direction: column;

}
.product-card {
}


</style>
    