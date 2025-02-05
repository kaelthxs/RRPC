<script setup>

</script>

<template>
  <header class="header">
    <div class="header__container">
      <div class="header__logo-wrapper">
        <div class="logo__wrapper-logoRR">
          RR
        </div>
        <div class="logo__wrapper-underlogoRR">
          Website to sell some Glebs for sex-dolls
        </div>
      </div>

      <div class="header__search-wrapper">
        <input type="text" placeholder="Поиск" class="search-wrapper__input">
      </div>
      <div class="header__login-wrapper">
        <button class="login-wrapper__button-login">Войти</button>
      </div>
    </div>
  </header>
</template>

<style scoped>

body {
  padding-left: 194px;
  padding-right: 194px;
}
.header {
  height: 123px;
  //width: 100%;
  background-color: #272727;
  //display: flex;
  //align-items: center;
  //justify-content: center;
  //padding: 0 25px 0 25px;
}
.header__container{
  max-width: 100%;
  max-height: 123px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.header__search-wrapper {
  height: 55px;
  width: 455px;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #323232;
}
.header__login-wrapper {
  margin-left: 250px;
  display: flex;
  align-items: center;
  justify-content: center;

}
.header__logo-wrapper {
  //margin-right: 10px;
  display: flex;
  max-width: 100%;
  align-items: center;
  flex-direction: column;
}
.search-wrapper__input {
  height: 55px;
  width: 455px;
  background-color: #323232;
  border: 0px;
}

.logo__wrapper-logoRR{
  color: white;
  font-size: 64px;
}


</style>