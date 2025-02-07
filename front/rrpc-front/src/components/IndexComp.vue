<template>
    <div class="main">
        <div class="container">
            <div class="main__card" v-for="(card, index) in cards" :key="card.id" :infoAboutCard="cards[index]">
                <img src="../../src/assets/imgs/VideoCarta.png" alt="" class="card__photo">
                <div class="card__undertext" >
                    {{ card.Name }}
                </div>
            </div>
        </div>
    </div>
</template>

<script>

import axios from 'axios';
export default {
  name: 'ArtistsPage',
  data() {
    return {
      cards: [
        {
          id: 1,
          image: '',
          Name: '',
        }
      ]
    }
  },
  methods: {
    async getAllCategories() {
        if (localStorage.getItem('authToken') != '') {
            const response = await axios.get('http://localhost:8080/api/category/', {
            headers: {
            'Authorization': `Bearer ${localStorage.getItem('authToken')}`
        }   
            });
            this.response = response.data;
            this.cards = response.data.map(item => ({
                id: item.id,
                image: item.img || 'image 1.png',
                Name: item.Name

            }));
        }
    }
  },
  mounted() {

    this.getAllCategories();
  }
}

</script>

<style scoped>
.main {
    background-color: #222222;
    width: 100%;
    height: calc(100vh - 123px);
}

.main__card {
    width: 250px;
    height: 250px;
    background-color: #272727;
    border-radius: 50px;
    margin-top: 50px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    transition: .5s;
}

.main__card:hover {
    transform: scale(1.1);

}

.container {
    padding-left: 150px;
    /* padding-right: 150px; */
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    align-items: center;
    width: 100%;
}
.card__photo {
    margin-top: 20px;
    width: 100%;
    height: 60%;

}
.card__undertext {
    background-color: #323232;
    width: 100%;
    height: 25%;
    border-bottom-left-radius: 200px;
    border-bottom-right-radius: 200px;
    border-top-right-radius: 100px;
    border-top-left-radius: 100px;
    justify-content: center;
    display: flex;
    align-items: center;
    color: white;
}

</style>
    