<template>
    <div class="auth__bgc" @click.self="closeAuthForm">
        <div class="auth">
            <div class="auth__title">
                <button class="title__button--login" @click="showAuthForm">
                    Вход
                </button>
                <button class="title__button--register" @click="showRegisterForm">
                    Регистрация
                </button>
            </div>

            <input type="text" name="" id="" placeholder="Логин" class="auth__input--login" v-model="dataUser.username" v-if="ChosenAuth">
            <input type="text" name="" id="" placeholder="Пароль" class="auth__input--password" v-model="dataUser.password_hash" v-if="ChosenAuth">

            <button class="auth__forgot--password--button" v-if="ChosenAuth">
                Забыли пароль?
            </button>

            <button class="auth__login--button" @click="sendSignIn" v-if="ChosenAuth">
                Войти
            </button>

            <input type="text" name="" id="" placeholder="Логин" class="auth__input--login" v-model="dataNewUser.Username" v-if="ChosenRegister">
            <input type="text" name="" id="" placeholder="Почта" class="auth__input--email" v-model="dataNewUser.Email" v-if="ChosenRegister">
            <input type="text" name="" id="" placeholder="Пароль" class="auth__input--password" v-model="dataNewUser.Password" v-if="ChosenRegister">
            <input type="text" name="" id="" placeholder="Повторите пароль" class="auth__input--password" v-model="dataNewUser.repeatPassword" v-if="ChosenRegister">

            <button class="auth__register--button" @click="checkForm" v-if="ChosenRegister">
                Зарегистрироваться
            </button>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'AuthComp',
  components: {

  },
  props: ['closeAuth'], // Получаем функцию закрытия из App.vue
  data () {
    
    return {
    JWTAvailability: false,
    ChosenAuth: true,
    ChosenRegister: false,
    roleOfAuthUser: 'admin',

    dataNewUser: {
        Username: '',
        Email: '',
        Password: '',
        repeatPassword: '',
        Role: 'admin'
    },
      dataUser: {
        username: '',
        password_hash: ''
      },
      token: localStorage.getItem('authToken'),
    }
  },
  methods: {

    sendRoleToApp() {
            this.$emit('RoleOfUser', this.dataNewUser.Role); // Передаём событие наверх
    },

    closeAuthForm() {
        this.closeAuth();
    },
    sendDataToBack () {
      console.log(this.dataUser)
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
            this.response = response.data;
            console.log(this.response)
            this.roleOfAuthUser = response.data['Role'];
            console.log(this.roleOfAuthUser)
        }
    },

    reloadPage() {
      window.location.reload();
    },

    async sendSignIn() {
      try {
        const response = await axios.post('http://localhost:8080/auth/sign-in', JSON.stringify(this.dataUser), { headers: { 'Content-Type': 'application/json' }
        });
        if (response.data.token) {
          localStorage.setItem('authToken', response.data.token)
          console.log('Токен сохранен в localStorage:', response.data.token)
          this.JWTAvailability = true
          this.$emit('JWTAvailability', true) // Передача события наверх
          this.getUserRoleById()
          this.$emit('RoleOfUser', this.roleOfAuthUser) // Передача роли пользователя
          this.closeAuthForm()
          this.reloadPage()
        }
        
        this.responseMessage = response.data.message;

      } catch (error) {
        console.error('Ошибка при отправке данных:', error);
        console.log(this.dataUser)
      }
    },

    async sendSignUp() {
      try {
        const response = await axios.post('http://localhost:8080/auth/sign-up', JSON.stringify(this.dataNewUser), { headers: { 'Content-Type': 'application/json' }
        });
        if (response.data.token) {
          localStorage.setItem('authToken', response.data.token);
          console.log('Токен сохранен в localStorage:', response.data.token);
          this.$router.push('/');
        }
        this.responseMessage = response.data.message;
        console.log(response.data.token);
        this.reloadPage()
      } catch (error) {
        console.error('Ошибка при отправке данных:', error);
        console.log(this.dataUser)
      }
    },
    checkForm () {
        console.log(this.dataNewUser);
        
      if (this.dataNewUser.Username && this.dataNewUser.Email && this.dataNewUser.Password && this.dataNewUser.repeatPassword) {
        if (this.dataNewUser.Password === this.dataNewUser.repeatPassword && this.dataNewUser.Password.length >= 6) {
            if (this.dataNewUser.Email.includes('@') && this.dataNewUser.Email.includes('.')) {
                return alert ('Вы успешно зарегистрировались ', this.dataNewUser.Username),
                this.sendSignUp()
            } else {
              return alert('Введите корректный email')
            }
        } else {
        return alert('Пароли не совпадают или меньше 6 символов')
        }
    } else {
        return alert('Введите все данные')
    }
},
    
    showAuthForm() {
        this.ChosenAuth = true;
        this.ChosenRegister = false;
    },
    showRegisterForm() {
        this.ChosenRegister = true;
        this.ChosenAuth = false;
    }
  }
}

</script>

<style scoped>

.title__button--login {
    background-color: #323232;
    color: white;
    border-radius: 50px 0 0 0;
    height: 70px;
    width: 100%;
    outline: none;
    border: none;
}
.title__button--login:hover {
    background-color: #5a5a5a;
}
.title__button--register {
    background-color: #323232;
    color: white;
    border-radius: 0 50px 0 0;
    height: 70px;
    width: 100%;
    outline: none;
    border: none;
}
.title__button--register:hover {
    background-color: #5a5a5a;
}

.auth__bgc {
    position: absolute;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.5);
}
.ifChosenAuth {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 300px;
    height: 200px;
    background-color: #272727;
    display: flex;
    flex-direction: column;
    border-radius: 50px;
    justify-content: space-between;
    align-items: center;
    z-index: 50;

}
.auth {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 300px;
    height: 450px;
    background-color: #272727;
    display: flex;
    flex-direction: column;
    border-radius: 50px;
    justify-content: space-between;
    align-items: center;
    z-index: 50;
}
.auth__title {
    display: grid;
    grid-template-columns: 1fr 1fr;
    background-color: #323232;
    color: white;
    border-radius: 50px 50px 0 0;
    height: 70px;
    width: 100%;
}
.auth__input--login {
    background-color: #222222;
    color: #555555;
    outline: none;
    border: none;
    padding-left: 20px;
    height: 50px;
    width: 80%;
    border-radius: 20px;
}
.auth__input--password {
    background-color: #222222;
    color: #555555;
    outline: none;
    border: none;
    padding-left: 20px;
    height: 50px;
    width: 80%;
    border-radius: 20px;
}
.auth__input--email {
    background-color: #222222;
    color: #555555;
    outline: none;
    border: none;
    padding-left: 20px;
    height: 50px;
    width: 80%;
    border-radius: 20px;
}
.auth__forgot--password--button {
    color: #5a5a5a;
    border: none;
    outline: none;
    background-color: transparent;
    transition: .3s;
}
.auth__forgot--password--button:hover {
    color: white;
}
.auth__login--button {
    margin-bottom: 20px;
    background-color: #222222;
    color: white;
    border: none;
    outline: none;
    width: 100px;
    height: 50px;
    border-radius: 10px;
    transition: .3s;
}
.auth__login--button:hover {
    background-color: #5a5a5a;
}

.auth__register--button {
    margin-bottom: 20px;
    background-color: #222222;
    color: white;
    border: none;
    outline: none;
    width: 150px;
    height: 50px;
    border-radius: 10px;
    transition: .3s;
}
.auth__register--button:hover {
    background-color: #5a5a5a;
}

</style>
    