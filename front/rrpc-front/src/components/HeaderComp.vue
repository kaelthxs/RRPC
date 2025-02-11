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
                    <button class="profile__menu--button-user">
                        Личный кабинет
                    </button>
                    <button class="profile__menu--button-cart">
                        Корзина
                    </button>
                    <button @click="logout" class="profile__menu--button-logout">
                        Выйти
                    </button>
                </div>
            </div>

        </div>
    </div>

</div>
<!-- <AuthComp></AuthComp> -->
</template>

<script>
// import AuthComp from './components/AuthComp.vue'
export default {
  name: 'HeaderComp',
  components: {
    // AuthComp
  },
  data () {
    return {
        visibility: false,
        showProfileMenu: false
  }
},
props: ['jwtAvailable', 'userRole'],

methods: {

    showProfileMenuFunc() {
        this.showProfileMenu = !this.showProfileMenu
    },

    logout() {
      localStorage.setItem('authToken', '');
      this.$emit('JWTAvailability', false);
      this.$emit('RoleOfUser', '');
      this.reloadPage()
    },

    reloadPage() {
      window.location.reload();
    },

    showForm() {
        this.visibility = !this.visibility;
        this.$emit('toggle-auth', this.visibility); // Передаём событие наверх
    }
}
}
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
</style>
