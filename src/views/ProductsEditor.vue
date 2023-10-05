<script setup>
</script>
<template>
    <div class="container mt-5">
        <div class="row">
            <div class="col-md-12">
                <h1>Список товаров</h1>

                <table class="table">
                    <thead>
                        <tr>
                            <th>Название</th>
                            <th>Описание</th>
                            <th>Цена</th>
                            <th>Действия</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(product, index) in poolProducts" :key="index">
                            <td>
                                <span v-if="!product.editing">{{ product.name }}</span>
                                <input v-else v-model="product.name" />
                            </td>
                            <td>
                                <span v-if="!product.editing">{{ product.description }}</span>
                                <input v-else v-model="product.description" />
                            </td>
                            <td>
                                <span v-if="!product.editing">{{ product.price }}</span>
                                <input v-else v-model="product.price" />
                            </td>
                            <td>
                                <button class="button-product-edit" @click="editProduct(index)" v-if="!product.editing"><i class="bi bi-pencil-fill"></i></button>
                                <button class="button-product-yes" @click="saveProduct(index)" v-if="product.editing"><i class="bi bi-check2"></i></button>
                                <button class="button-product-no" @click="discardProductChanges(index)" v-if="product.editing"><i class="bi bi-x-lg"></i></button>
                                <button class="button-product-delete" @click="deleteProduct(index)"><i class="bi bi-trash3-fill"></i></button>
                            </td>
                        </tr>
                    </tbody>
                </table>

                <div class="button-add">
                    <button @click="addProduct()" class="styled-button">+ Добавить товар</button>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
    export default {
        data() {
            return {
                poolProducts: [],
                isLoggedIn: false,
                name: "",
                rules: ""
            };
        },
        async mounted() {
            this.loadListProducts();

            try {
				await this.checkAuth();
                
				if (this.isLoggedIn) {
                    if(this.rules == "user") {
                        window.location.href = '/';
                    }
				    
				}
			} catch (error) {
				console.error(error);
			}
        },
        methods: {
            async checkAuth() {
				const response = await fetch(`/check_auth`, {
					method: 'GET',
				});

				if (!response.ok) {
					throw new Error('Не удалось получить данные о авторизации');
				}

				const data = await response.json();

				if (data.isLoggedIn) {
					this.isLoggedIn = true;
					this.name = data.name;
                    this.rules = data.rules;
				} else {
					this.isLoggedIn = false;
					this.name = "";
				}
			},
            async loadListProducts() {
                try {
                    const response = await fetch(`/get_all_products`, {
                        method: 'GET',
                    });

                    if (!response.ok) {
                        throw new Error('Не удалось получить данные о продуктах');
                    }

                    const data = await response.json();
                    const formattedData = data.map((product) => ({
                        name: product.name,
                        description: product.description,
                        price: product.price,
                    }));

                    this.poolProducts = formattedData;
                } catch (error) {
                    console.error(error);
                }
            },

            saveProduct(index) {
                this.editProductOnServer({
                    index,
                    name: this.poolProducts[index].name,
                    description: this.poolProducts[index].description,
                    price: this.poolProducts[index].price,
                }).then(() => {
                    this.poolProducts[index].editing = false;
                }).catch((error) => {
                    alert('Ошибка при сохранении продукта на сервере');
                });
            },

            discardProductChanges(index) {
                this.poolProducts[index].name = this.selectedProductContent.name;
                this.poolProducts[index].description = this.selectedProductContent.description;
                this.poolProducts[index].price = this.selectedProductContent.price;
                this.poolProducts[index].editing = false;
            },

            addProduct() {
                const newProduct = {
                    name: "Новый продукт",
                    description: "",
                    price: 0,
                };

                this.addProductOnServer(newProduct).then(() => {
                    this.poolProducts.push(newProduct);
                }).catch((error) => {
                    alert('Ошибка при сохранении продукта на сервере:', error);
                });
            },

            editProduct(index) {
                this.selectedProductContent = {
                    name: this.poolProducts[index].name,
                    description: this.poolProducts[index].description,
                    price: this.poolProducts[index].price,
                };

                this.poolProducts[index].editing = true;
            },

            deleteProduct(index) {
                const confirmDelete = confirm('Вы уверены, что хотите удалить этот продукт?');
                if (confirmDelete) {
                    this.poolProducts.splice(index, 1);
                    this.deleteProductOnServer({ index });
                }
            },

            async addProductOnServer(product) {
                try {
                    const response = await fetch('/add_product', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json',
                        },
                        body: JSON.stringify(product),
                    });

                    if (!response.ok) {
                        throw new Error('Не удалось сохранить продукт на сервере');
                    }

                    return true;
                } catch (error) {
                    console.error('Ошибка при отправке запроса на сервер:', error);
                    throw error;
                }
            },

            async deleteProductOnServer(product) {
                try {
                    const response = await fetch('/delete_product', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/x-www-form-urlencoded',
                        },
                        body: new URLSearchParams({
                            index: product.index + 1,
                        }),
                    });
                } catch (error) {
                    console.error('Ошибка при отправке запроса на сервер:', error);
                }
            },

            async editProductOnServer(product) {
                try {
                    const response = await fetch('/edit_product', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/x-www-form-urlencoded',
                        },
                        body: new URLSearchParams({
                            index: product.index,
                            name: product.name,
                            description: product.description,
                            price: product.price,
                        }),
                    });
                } catch (error) {
                    console.error('Ошибка при отправке запроса на сервер:', error);
                }
            },
        },
    };
</script>
<style>
    .container {
        margin-top: 50px;
    }

    h1 {
        font-size: 24px;
        margin-bottom: 20px;
    }

    table {
        width: 100%;
        border-collapse: collapse;
        border: 1px solid #ddd;
        border-radius: 5px;
    }

    th,
    td {
        border: 1px solid #ddd;
        padding: 8px;
        text-align: left;
    }

    th {
        background-color: #573232;
        border: 1px solid #ddd;
    }

    .button-add {
        margin-top: 20px;
    }

    .styled-button {
        background-color: #202020;
        color: #fff;
        border: none;
        padding: 10px 20px;
        border-radius: 5px;
        cursor: pointer;
        display: block;
        width: 100%;
        text-align: center;
    }

    .button-product-edit,
    .button-product-yes,
    .button-product-no,
    .button-product-delete {
        color: #fff;
        border: none;
        padding: 5px 10px;
        border-radius: 5px;
        cursor: pointer;
        margin-right: 5px;
    }

    .button-product-edit {
        background-color: #10a352;
    }

    .button-product-yes {
        background-color: #10a352;
    }

    .button-product-no {
        background-color: #a71111;
    }

    .button-product-delete {
        background-color: #a71111;
    }

    input {
        width: 100%;
        padding: 5px;
        border: 1px solid #ddd;
        border-radius: 5px;
    }
</style>