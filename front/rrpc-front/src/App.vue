<template>
  <HeaderComp/>
  <IndexComp/>
  <AuthComp/>
</template>

<script>
import HeaderComp from './components/HeaderComp.vue'
import IndexComp from './components/IndexComp.vue'
import AuthComp from './components/AuthComp.vue'

export default {
  name: 'App',
  components: {
    HeaderComp,
    IndexComp,
    AuthComp
  }
}



</script>

<style>
  body{margin: 0;padding: 0;box-sizing: border-box;max-width: 100%; font-family: "Roboto", serif;} 
  h1, h2, h3, h4, h5, h6, ul, ol, li, p{margin: 0;padding: 0;} 
  ul, ol{list-style: none;} 
  a{text-decoration: none;color: black;} 
  img {max-width: 100%;display: block;}

</style>
