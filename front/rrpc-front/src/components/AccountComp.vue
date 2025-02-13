<template>
    <div class="main">
        <div class="container">
            <div class="main__AccountInfo">
                <span>Имя: {{ cards.Username }}</span>
                <span>Почта: {{ cards.Email }}</span>
                <span>Пароль: *******</span>
            </div>
            <button @click="$emit('back')">Назад</button>
            <div class="main__orders">
                <h2>Мои заказы</h2>
                <ul>
                    <li v-for="order in orders" :key="order.ID">
                    Заказ №{{ order.ID }} — {{ order.TotalPrice }}₽
                    </li>
                </ul>
            </div>

        </div>
    </div>
</template>

<script>

import axios from 'axios';
export default {
  name: 'AccountComp',
  data() {
    return {
      cards: [],
      orders: []
    }
  },
  methods: {

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
            const response = await axios.get(`http://localhost:8080/api/orders/${userId}`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }   
            });

            console.log("Ответ от сервера:", response.data); // Дебаг
            this.orders = response.data; // Если массив, то оставляем как есть
            console.log("заказы");
            console.log(this.orders);

            
            
        } catch (error) {
            console.error("Ошибка при получении данных пользователя:", error);
        }
    }


  },
  mounted() {
    console.log("Рендерим AccountComp")
    this.getUserById()
    this.getOrdersByUserId()
  }
}

</script>

<style scoped>

.main {
    background-color: #222222;
    width: 100%;
    height: calc(100vh - 123px);
}

.main__orders {
    color: white;
}
.container {
    align-items: center;
    justify-content: center;
    width: 80%;
    justify-items: center;
}
.main__AccountInfo {
    color: white;
    display: flex;
    flex-direction: column;
    height: 300px;
}

</style>
    