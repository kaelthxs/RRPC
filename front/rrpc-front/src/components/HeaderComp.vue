<template>
    <div class="header">
      <div class="container">
        <div class="header__logo-wrapper" @click="$emit('back')">
          <div class="logo-wrapper__RR">RR</div>
          <div class="logo-wrapper__underRR"></div>
        </div>
        <div class="header__input-wrapper">
          <input
            type="text"
            class="header__input"
            placeholder="Поиск"
            v-model="searchQuery"
            @input="searchProducts"
            @focus="showResults = true"
          />
          <button class="header__input-logo--wrapper">
            <svg class="header__input-logo">
              <use href="../../src/assets/sprite.svg#search"></use>
            </svg>
          </button>
        </div>

        <div v-if="showResults && filteredProducts.length" class="search-results">
            <div v-for="product in filteredProducts" :key="product.id" class="search-result-item" @click="selectProduct(product)">
              {{ product.Creator }} {{ product.Name }}
              <!-- <button @click="addToCart(product.ID)">Добавить в корзину</button> -->

            </div>
          </div>
  
        <div class="header__login-wrapper">
          <button class="login-wrapper__button" @click="showForm" v-if="!jwtAvailable">
            Войти
          </button>
  
          <div class="header__profile" v-if="jwtAvailable">
            <button class="profile__svg--button" @click="showProfileMenuFunc">
              <svg class="profile__svg">
                <use href="../../src/assets/sprite.svg#user"></use>
              </svg>
            </button>
            <div class="profile__menu" v-if="showProfileMenu">
              <button class="profile__menu--button-user" @click="navigateToAccount">Личный кабинет</button>
              <button class="profile__menu--button-cart" @click="navigateToCart">Корзина</button>
              <button @click="logout" class="profile__menu--button-logout">Выйти</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import axios from "axios";
  
  export default {
    name: "HeaderComp",
    props: ["jwtAvailable", "userRole"],
    data() {
      return {
        visibility: false,
        showProfileMenu: false,
        searchQuery: '',
        showResults: false,
        allProducts: [],
        cart: []
      };
    },
    computed: {
        filteredProducts() {
            if (!this.searchQuery.trim()) return [];
            return this.allProducts.filter(product => 
            product.Name  && product.Name.toLowerCase().includes(this.searchQuery.toLowerCase())
            );
        }
    },
    methods: {
        async getAllProducts() {
            console.log("getAllProduct вызван");
            

        try {
            const authToken = localStorage.getItem('authToken');
            if (!authToken) {
                console.warn("Нет токена авторизации.");
                return;
            }

            const response = await axios.get('http://localhost:8080/api/product/', {
                headers: {
                    'Authorization': `Bearer ${authToken}`
                }
            });

            console.log("Ответ сервера:", response.data);
            
            this.allProducts = response.data;

            // Доступ к массиву
            const productsArray = response.data; // Проверка на data в ответе
            console.log("Массив продуктов:", productsArray);

            if (Array.isArray(productsArray)) {
                this.products = productsArray.map(item => ({
                    ID: item['ID'],
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

    // addToCart (id) {
    //     console.log(id);
    //     console.log(this.products);
    //     for (const el of this.products) {
    //         if (el['id'] == id) {
    //             this.cart.push(el)
    //         }
    //     }
    //     console.log('ВЫВОД КОРЗИНЫ НЕ ИЗ LS');
        
    //     console.log(this.cart);
    //     console.log(JSON.stringify(this.cart));
        
    //     console.log(`----------`);

    //     localStorage.setItem('cart', JSON.stringify(this.cart))

    //     console.log(`------------`);
    // },

      searchProducts() {
        this.showResults = true;
      },
  
      selectProduct(product) {
        this.searchQuery = product.Name;
        this.showResults = false;
      },
  
      closeSearchResults(event) {
        if (!this.$el.contains(event.target)) {
          this.showResults = false;
        }
      },
  
      navigateToAccount() {
        this.$emit("navigate", "account");
      },
  
      navigateToCart() {
        this.$emit("navigate", "cart");
      },
  
      showProfileMenuFunc() {
        this.showProfileMenu = !this.showProfileMenu;
      },
  
      logout() {
        localStorage.setItem("authToken", "");
        this.$emit("JWTAvailability", false);
        this.$emit("RoleOfUser", "");
        this.reloadPage();
      },
  
      reloadPage() {
        window.location.reload();
      },
  
      showForm() {
        this.visibility = !this.visibility;
        this.$emit("toggle-auth", this.visibility);
      },
    },
    mounted() {
        this.getAllProducts();
        document.addEventListener("click", this.closeSearchResults);

        // Загружаем корзину из локального хранилища
        const savedCart = JSON.parse(localStorage.getItem('cart')) || [];
        this.cart = savedCart;
    },
        beforeUnmount() {  // Исправленный хук
        document.removeEventListener("click", this.closeSearchResults);
    }

  };
  </script>

<style scoped>

.header__profile {
    display: flex;
    justify-content: flex-end;

}
.profile__svg--button {
    background-color: transparent;
    border: none;
    outline: none;
    fill: white;
    height: 46px;
    width: 40px;
    transition: .3s;
}

.profile__svg--button:hover {
    transform: scale(1.1);
}
.profile__svg {
    height: 46px;
    width: 40px;

}
.profile__menu {
    position: absolute;
    margin-top: 60px;
    display: flex;
    flex-direction: column;
    background-color: black;
    width: 300px;
    right: 50%;
    transform: translateX(50%);
    border-radius: 50px;
}

.profile__menu--button-user {
    background-color:#323232;
    outline: none;
    border: none;
    color: white;
    height: 50px;
    border-radius: 40px 40px 0px 0px;
}
.profile__menu--button-cart {
    background-color:#323232;
    outline: none;
    border: none;
    color: white;
    height: 50px;

}
.profile__menu--button-logout {
    background-color:#323232;
    outline: none;
    border: none;
    color: white;
    height: 50px;
    border-radius: 0px 0px 40px 40px;
}

.profile__menu--button-user:hover {
    background-color:#5a5a5a;
}

.profile__menu--button-cart:hover {
    background-color:#5a5a5a;
}

.profile__menu--button-logout:hover {
    background-color:#5a5a5a;
}


.header__profile {
    position: relative;
}

.header {
    width: 100%;
    height: 123px;
    background-color: #272727;
    display: flex;
    flex-direction: row;
}
.container {
    padding-left: 150px;
    padding-right: 150px;
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    align-items: center;
    width: 100%;
}
.header__logo-wrapper {
    display: flex;
    flex-direction: column;
    pointer-events: none;
}

.header__input-wrapper {
    display: flex;
    align-items: center;
    width: 100%;
    background-color: #323232;
    border-radius: 10px;

}
.header__input {
    border: none;
    background-color: #323232;
    height: 40px;
    width: 100%;
    outline: none;
    border-radius: 10px;
    padding-left: 20px;

}

.header__input-logo {
    fill: #5a5a5a;
    width: 22px;
    height: 22px;
    transition: .3s;
}

.header__input-logo:hover {
    fill: white;
}

.header__input-logo--wrapper {
    width: 0%;
    height: 0%;
    background-color: transparent;
    border: none;
    margin-right: 25px;
}

.header__login-wrapper {
    display: flex;
    justify-content: flex-end;

}

.login-wrapper__button {
    background-color: transparent;
    border: none;
    color: white;
    font-size: 18px;
    justify-content: right;
}

.logo-wrapper__RR {
    font-family: "Sarina", serif;
    color: white;
    font-size: 50px;
    margin-left: 88px;
    pointer-events: none;
}

.logo-wrapper__underRR {
    font-family: "Roboto", serif;
    color: white;
    pointer-events: none;
}

.search-results {
  position: absolute;
  background:#5a5a5a;
  margin-top: 200px;
  margin-left: 420px;
  width: 300px;
  max-height: 100px;
  overflow-y: auto;
  z-index: 100;
  display: flex;
  flex-direction: column;
  border-radius: 30px;
}

.search-result-item {
    background-color: #2B2B2B;
    margin-top: 5px;
    height: 75px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 10px;
    color: white;
    transition: .3s;
}
</style>
