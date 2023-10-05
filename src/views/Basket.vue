<script setup>
</script>
<template>
    <div class="container mt-5">
        <h1 class="text-center mb-4">Корзина</h1>

        <div class="my-box">
            <div class="my-box-box" v-for="(item, index) in basket" :key="index">
                <div class="row">
                    <div class="col-md-2">
                        <img :src="item.photoUrl" class="product-image" alt="Product Image">
                    </div>
                    <div class="col-md-10 d-flex flex-column">
                        <div class="product-name">{{ item.name }}</div>
                        <div class="product-description">{{ item.description }}</div>
                        <div class="flex-grow-1"></div>
                        <div class="d-flex justify-content-between align-items-center">
                            <div class="product-price">Цена: <strong>{{ item.price }}</strong></div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            basket: []
        };
    },
    mounted() {
        this.getBasket();
    },
    methods: {
        async getBasket() {
            const response = await fetch(`/get_basket_user`, {
                method: 'GET',
            });

            if (!response.ok) {
                throw new Error('Не удалось получить данные о авторизации');
            }

            const data = await response.json();

            this.basket = data
        }
    }
}
</script>
<style>
.my-box {
    background-color: #f1f1f1;
    border-radius: 10px;
    padding: 20px;
    margin-top: 20px;
}

.my-box-box {
    background-color: #b6b6b662;
    border-radius: 10px;
    padding: 20px;
    margin-top: 20px;
}
</style>