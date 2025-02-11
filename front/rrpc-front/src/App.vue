<template>
  <HeaderComp :jwtAvailable="jwtAvailable" :userRole="userRole" @toggle-auth="toggleAuth" />
  <AuthComp v-if="isAuthVisible" :closeAuth="closeAuth" 
      @JWTAvailability="updateJWTAvailability"
      @RoleOfUser="updateUserRole"
      @closeAuth="showAuth = false"/>
      <IndexComp v-if="currentView === 'index'" @open-search="openSearch"/>
      <ProductsComp v-if="currentView === 'search'" :categoryId="selectedCategoryId" @back="goBack" />
</template>

<script>
import HeaderComp from './components/HeaderComp.vue';
import AuthComp from './components/AuthComp.vue';
import IndexComp from './components/IndexComp.vue';
import ProductsComp from './components/ProductsComp.vue'

export default {
  components: {
    HeaderComp,
    AuthComp,
    IndexComp,
    ProductsComp
  },
  data() {
    return {
      isAuthVisible: false,
      jwtAvailable: !!localStorage.getItem('authToken'),
      userRole: '',
      currentView: 'index',
      selectedCategoryId: null
    };
  },
  methods: {

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
    }

    
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
