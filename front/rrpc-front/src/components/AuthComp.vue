<template>
    <div class="auth">
        <div class="auth__title">
            Вход
        </div>
        <input type="text" name="" id="" placeholder="Логин" class="auth__input--login">
        <input type="text" name="" id="" placeholder="Пароль" class="auth__input--password">
        <button class="auth__forgot--password--button">
            Забыли пароль?
        </button>
        <button class="auth__login--button">
            Вход
        </button>
    </div>
</template>

<script>

</script>

<style scoped>

.auth {
    width: 300px;
    height: 350px;
    background-color: #272727;
    display: flex;
    flex-direction: column;
    border-radius: 50px;
    justify-content: space-between;
    align-items: center;
}
.auth__title {
    background-color: #323232;
    color: white;
    border-radius: 50px 50px 0 0;
    height: 70px;
    display: flex;
    justify-content: center;
    align-items: center;
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

</style>
    