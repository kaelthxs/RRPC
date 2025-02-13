<template>
    <button @click="$emit('back')">Назад</button>
    <div class="main__main">
        <div>
            <ul>
                <li v-for="(product, index) in products" :key="product[index]" class="cart__spisok">
                    <div class="product__item">
                        <div class="item__title">{{ product.Name }} {{ product.Creator }} </div> 
                        <div class="item__price"> Цена:{{ product.Price }} ₽</div>
                    </div>
                </li>
            </ul>
            <div class="final_price">Финальная цена: {{ totalPrice }} ₽</div>
            <button @click="createNewOrder">Создать заказ</button>
            <button @click="clearCart">Очистить корзину</button>
            <button @click="goBack">Назад</button>
        </div>
    </div>
</template>

<script>

import axios from 'axios';
export default {
    props: {
    cart: Array
  },
    name: 'CartComp',
    data() {
        return {
            localCart: [...this.cart],  // Копируем корзину в локальный массив
            products: []
        }
    },

    computed: {
        cartItems() {
            console.log("Корзина в CartComp:", this.cart);
            return this.cart;
        },

        totalPrice() {
            let sum = 0
            console.log(this.products);
            
            if (localStorage.getItem('cart') !== null){
                this.products.forEach(element => {
                    sum += element['Price']
                })
            }
            return sum;
        }
    },

    watch: {
        cart: {
            handler(newCart) {
                this.$emit('update-cart', newCart);
            },
            deep: true
        }
    },
  methods: {

    updateQuantity(id, change) {
        // this.products[0]['Quantity'] = 1124124
        for (const key in this.products) {
            const el = this.products[key]; // Получаем объект продукта

            if (el['ID'] == id) { 
                if (Object.hasOwn(el, "Quantity")) {
                    el['Quantity'] += 1;
                } else {
                    // Vue может не отследить добавление нового свойства, поэтому явно обновляем его
                    this.$set(el, "Quantity", 1);  // В Vue 2
                    // el.Quantity = 1; // В Vue 3 можно просто так
                    this.$forceUpdate(); // Принудительное обновление Vue 3
                }
            }
        }

        console.log(this.products);
        
      this.$emit('update-quantity', { id, change });
    },
    removeFromCart(id) {
        console.log(1231231213);
        console.log(id);
        this.products = this.products.filter(product => product.id !== id);
        // this.$emit('remove-from-cart', id);
    },
    clearCart() {
        localStorage.removeItem('cart')
        this.products = [];
    },

    saveCart() {
        const userId = this.getUserIdFromJWT();
        if (!userId) {
        console.warn("Не удалось сохранить корзину: нет userId");
        return;
        }

        localStorage.setItem(`cart_${userId}`, JSON.stringify(this.cart));
        console.log("Корзина сохранена в localStorage:", this.cart);
    },
 
    // async createOrder() {

    //     const token = localStorage.getItem('authToken');
    //     if (!token) {
    //         console.warn("JWT токен отсутствует");
    //         return;
    //     }

    //     const userId = this.getUserIdFromJWT(); 
    //   if (this.localCart.length === 0) {
    //     alert("Корзина пустая!");
    //     return;
    //   }

      

    //   try {
    //     const response = await axios.post(`https://localhost:8080/api/orders/addOrderByUserId/${userId}`, {
    //       items: this.localCart
    //     });

    //     console.log("Заказ успешно создан:", response.data);

    //     // Здесь, если заказ успешно оформлен, очищаем корзину
    //     this.localCart = [];
    //     // Можно также отправить событие, чтобы очистить корзину на родительском уровне
    //     this.$emit('update-cart', this.localCart);
    //   } catch (error) {
    //     console.error("Ошибка при оформлении заказа:", error);
    //   }
    // },

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

    async getUserById() {
        const token = localStorage.getItem('authToken');
        if (!token) {
            console.warn("JWT токен отсутствует");
            return;
        }

        const userId = this.getUserIdFromJWT(); 
        if (!userId) {
            console.warn("Не удалось получить userId из JWT");
            return;
        }

        try {
            const response = await axios.get(`http://localhost:8080/api/user/${userId}`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }   
            });

            console.log("Ответ от сервера:", response.data); // Дебаг
            this.cards = response.data; // Если массив, то оставляем как есть
            
        } catch (error) {
            console.error("Ошибка при получении данных пользователя:", error);
        }
    },

    async getOrdersByUserId() {
        const token = localStorage.getItem('authToken');
        if (!token) {
            console.warn("JWT токен отсутствует");
            return;
        }

        const userId = this.getUserIdFromJWT(); 
        if (!userId) {
            console.warn("Не удалось получить userId из JWT");
            return;
        }

        try {
            const response = await axios.get(`http://localhost:8080/api/user/${userId}`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }   
            });

            console.log("Ответ от сервера:", response.data); // Дебаг
            this.cards = response.data; // Если массив, то оставляем как есть
            
        } catch (error) {
            console.error("Ошибка при получении данных пользователя:", error);
        }
    },
    showCartItems (){
        const cart = JSON.parse(localStorage.getItem('cart'))
        this.products = cart
        console.log(this.products);
    },

    async createNewOrder() {
    try {
        const UserID = this.getUserIdFromJWT();
        const token = localStorage.getItem('authToken');
        const totalPrice = this.totalPrice;

        // 1. Создаём заказ и получаем его ID
        const orderResponse = await axios.post(
            `http://localhost:8080/api/orders/addOrderByUserId/${UserID}`, 
            { UserID, totalPrice }, 
            { 
                headers: { 
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                }   
            }
        );
        console.log("=-------------------=");
        console.log(" АВЫОАВЫОШАВЫ"+orderResponse);
        console.log("=-------------------=");

        

        const orderId = orderResponse.data.id; // Получаем ID заказа из ответа
        console.log("Заказ создан, ID:", orderId);
        console.log("Ответ сервера на создание заказа:", orderResponse.data);


        // 2. Добавляем товары в заказ
        const cart = JSON.parse(localStorage.getItem('cart')) || [];
        console.log(cart);

        

        for (const el of cart) {

            console.log("Отправка товара:", {
                order_id: orderId, 
                product_id: el.ID, 
            });

            await axios.post(
                `http://localhost:8080/api/orderItem/${orderId}`, 
                { 
                    OrderID: orderId, 
                    ProductID: el.ID, 
                },
                { 
                    headers: { 
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`
                    }   
                }
            ); 
            console.log(`Товар ${el.ID} добавлен в заказ ${orderId}`);
        }

        // 3. Очищаем корзину и обновляем интерфейс
        this.localCart = [];
        localStorage.removeItem('cart');
        this.$emit('update-cart', this.localCart);
        alert(`Заказ №${orderId} успешно создан!`);

    } catch (error) {
        console.error('Ошибка при создании заказа:', error);
    }
}


},
  mounted() {
    this.showCartItems()
    this.getUserById()
    this.getOrdersByUserId()
    console.log("Текущая корзина в CartComp:", this.cart);
  }
}

</script>

<style scoped>
.main__main {
    display: flex;
    flex-direction: column;
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

.final_price {
    color: white;
}
</style>
    