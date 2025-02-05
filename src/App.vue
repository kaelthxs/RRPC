<template>
  <MainPage>
    Sup
  </MainPage>
  <HeaderComp>
    Sup2
  </HeaderComp>
</template>

<script>
import MainPage from './components/MainPage.vue'
import HeaderComp from "./components/HeaderComp.vue";

export default {
  name: 'IndexPage',
  components: {
    MainPage,
    HeaderComp,
  }
}
</script>

<style>
  body{margin: 0;padding: 0;box-sizing: border-box;max-width: 100%;}
  h1, h2, h3, h4, h5, h6, ul, ol, li, p{margin: 0;padding: 0;}
  ul, ol{list-style: none;}
  a{text-decoration: none;color: black;}
  img {max-width: 100%;display: block;}

</style>
