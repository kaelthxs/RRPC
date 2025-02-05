<template>
<div class="header">
    <div class="container">
        <div class="header__logo-wrapper">
            <div class="logo-wrapper__RR">
                RR
            </div>
            <div class="logo-wrapper__underRR">
                Web-site for selling Glebs to sex-dolls
            </div>
        </div>
        <div class="header__input-wrapper">
            <input type="text" class="header__input" placeholder="Поиск">
            <button class="header__input-logo--wrapper">
                <svg class="header__input-logo">
                    <use href="../../src/assets/sprite.svg#search"></use>
                </svg>
            </button>
        </div>
        <div class="header__login-wrapper">
            <button class="login-wrapper__button" @click="showForm">
                Войти
            </button>
        </div>
    </div>
</div>
</template>

<script>
// import AuthComp from './components/AuthComp.vue'
// export default {
//   name: 'AuthPage',
//   components: {
//     AuthComp
//   },
//   data () {
//     return {
//         visibility: false
//   }
// },
// methods: {
//     showForm() {
//         if (this.visibility == false) {
//             this.visibility = true 
//         } else {
//             this.visibility = false
//         }
//     }
// }
// }
</script>

<style scoped>
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
</style>
