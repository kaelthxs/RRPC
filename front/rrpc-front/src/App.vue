<template>
  <HeaderComp :jwtAvailable="jwtAvailable" :userRole="userRole" @toggle-auth="toggleAuth" @navigate="navigateTo" @back="goBack"/>
  <AuthComp v-if="isAuthVisible" :closeAuth="closeAuth" 
      @JWTAvailability="updateJWTAvailability"
      @RoleOfUser="updateUserRole"
      @closeAuth="showAuth = false"/>
      <IndexComp v-if="currentView === 'index'" @open-search="openSearch"/>
      <ProductsComp v-if="currentView === 'search'" :categoryId="selectedCategoryId" @back="goBack" @add-to-cart="addToCart" @update-cart="updateCart"/>
      <AccountComp v-if="currentView === 'account'" @back="goBack"/>
      <CartComp :cart="cart" v-if="currentView === 'cart'" @update-cart="updateCart" @back="goBack" @update-quantity="updateQuantity" @remove-from-cart="removeFromCart" @clear-cart="clearCart" />
</template>

<script>
import HeaderComp from './components/HeaderComp.vue';
import AuthComp from './components/AuthComp.vue';
import IndexComp from './components/IndexComp.vue';
import ProductsComp from './components/ProductsComp.vue'
import AccountComp from './components/AccountComp.vue'
import CartComp from './components/CartComp.vue'


export default {
  components: {
    HeaderComp,
    AuthComp,
    IndexComp,
    ProductsComp,
    AccountComp,
    CartComp
  },
  data() {
    return {
      cart: [],
      isAuthVisible: false,
      jwtAvailable: !!localStorage.getItem('authToken'),
      userRole: '',
      currentView: 'index',
      selectedCategoryId: null
    };
  },
  methods: {
    updateCart(){
      
    },
    updateQuantity({ id, change }) {
      const product = this.cart.find(item => item.id === id);
      if (product) {
        product.quantity = Math.max(1, product.quantity + change);
        this.saveCart();
      }
    },

    getUserIdFromJWT() {
      const token = localStorage.getItem('authToken'); 
      if (!token) return null;

      try {
        const payloadBase64 = token.split('.')[1]; // Получаем payload (вторая часть JWT)
        const decodedPayload = JSON.parse(atob(payloadBase64)); // Декодируем из base64
        return decodedPayload.UserId;
      } catch (error) {
        console.error('Ошибка декодирования JWT:', error);
        return null;
      }
    },


    loadCart() {
      const userId = this.getUserIdFromJWT();
      if (!userId) {
        console.warn("Не удалось загрузить корзину: нет userId");
        return;
      }

      const storedCart = localStorage.getItem(`cart_${userId}`);
      if (storedCart) {
        this.cart = JSON.parse(storedCart);
        console.log("Загружена корзина из localStorage:", this.cart);
      } else {
        console.log("Корзина пуста.");
      }
    },

    saveCart() {
      const userId = this.getUserIdFromJWT();
      if (!userId) {
        console.warn("Не удалось сохранить корзину: нет userId");
        return;
      }

      localStorage.setItem(`cart_${userId}`, JSON.stringify(this.cart)); // Используем userId в ключе
      console.log("Корзина сохранена в localStorage:", this.cart);
    },

  // Добавление товара в корзину
    addToCart(product) {
      console.log("Текущая корзина до добавления:", this.cart);
      
      const existingProduct = this.cart.find(item => item.id === product.id);
      if (existingProduct) {
        existingProduct.quantity += 1;
      } else {
        this.cart.push({ ...product, quantity: 1 });
      }

      this.saveCart();
      console.log("Корзина после добавления:", this.cart);
    },

    removeFromCart(productId) {
      this.cart = this.cart.filter(item => item.id !== productId);
      this.saveCart();
    },

    // Очистка корзины (например, при выходе из аккаунта)
    clearCart() {
      this.cart = [];
      const userId = this.getUserIdFromJWT();
      if (userId) {
        localStorage.removeItem(`cart_${userId}`);
      }
      console.log("Корзина очищена");
    },

    updateJWTAvailability(status) {
      this.jwtAvailable = status;
    },
    updateUserRole(role) {
      this.userRole = role;
    },

    toggleAuth(value) {
      this.isAuthVisible = value; // Обновляем состояние формы авторизации      
    },
    closeAuth() {
      this.isAuthVisible = false; // Закрываем окно
    },
    
    openSearch(categoryId) {
      this.selectedCategoryId = categoryId; 
      this.currentView = 'search';
      console.log("Переданный categoryId:", this.selectedCategoryId); // Дебаг
    },
    goBack() {
      this.currentView = 'index';
    },

    navigateTo(view) { 
        console.log("Получено событие navigate, переключаемся на:", view);
        this.currentView = view;
    },

    goCart() {
      this.currentView = 'cart';
    }

  },

  mounted () {
    this.loadCart();
  }
};
</script>

<style>
  body{margin: 0;padding: 0;box-sizing: border-box;max-width: 100%; font-family: "Roboto", serif; background-color: #222222;} 
  h1, h2, h3, h4, h5, h6, ul, ol, li, p{margin: 0;padding: 0;} 
  ul, ol{list-style: none;} 
  a{text-decoration: none;color: black;} 
  img {max-width: 100%;display: block;}

</style>
