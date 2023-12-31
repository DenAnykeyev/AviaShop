<script setup>
</script>

<template>
	<div>
		<h1 class="text-center">Наши продукты</h1>

		<div class="card shadow rounded my-3">
			<div class="pagination">
				<button v-if="currentPage > 1" @click="changePage(currentPage - 1)" class="page-button">
					Предыдущая
				</button>
				<template v-for="page in totalPages">
					<button @click="changePage(page)"
						:class="{ 'page-button': true, 'active-page': page === currentPage }">{{ page }}</button>
				</template>
				<button v-if="currentPage < totalPages" @click="changePage(currentPage + 1)" class="page-button">
					Следующая
				</button>
			</div>
			<div class="card shadow rounded my-3">
			</div>
			<div v-if="products.length > 0">
				<div v-for="product in products" :key="product.name">
					<div class="card-body">
						<div class="row">
							<div class="col-md-2">
								<img :src="product.photoUrl" class="product-image" alt="Product Image">
							</div>
							<div class="col-md-10 d-flex flex-column">
								<div class="product-name">{{ product.name }}</div>
								<div class="product-description">{{ product.description }}</div>
								<div class="flex-grow-1"></div>
								<div class="d-flex justify-content-end">
									<div class="product-price">Цена: <strong>{{ product.price }}</strong></div>
									<button class="btn btn-warning" title="В корзину" @click="addToBasket(product.id)">
										<i class="bi bi-cart"></i> <span class="d-none d-md-inline">В корзину</span>
									</button>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
			<p v-else>Нет доступных товаров</p>
		</div>
	</div>
</template>
<script>
export default {
	data() {
		return {
			// Для view
			products: [],
			currentPage: 1,
			totalPages: 0,

			// Для logic
			isLoggedIn: false,
			id: "",
			name: ""
		};
	},
	async mounted() {
		this.getProducts(1);
		this.getTotalPages();

		this.checkAuth();
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
			}
		},
		async changePage(page) {
			this.currentPage = page;
			await this.getProducts(page);
		},
		async getTotalPages() {
			try {
				const response = await fetch('/get_total_pages', {
					method: 'GET',
				});
				if (!response.ok) {
					throw new Error('Не удалось получить общее количество страниц.');
				}
				this.totalPages = await response.json();
			} catch (error) {
				console.error(error);
			}
		},
		async getProducts(page) {
			try {
				const response = await fetch(`/get_products?page=${page}`, {
					method: 'GET',
				});

				if (!response.ok) {
					throw new Error('Не удалось получить данные о продуктах');
				}

				const data = await response.json();
				const formattedData = data.map((product) => ({
					id: product.id,
					name: product.name,
					description: product.description,
					price: product.price,
					photoUrl: product.photoUrl
				}));

				this.products = formattedData;
			} catch (error) {
				console.error(error);
			}
		},
		async addToBasket(productId) {
			if (!this.isLoggedIn) {
				alert("Вы не вошли в аккаунт!")

				return;
			}

			const partBasket = {
				name: this.name,
				productId: productId
			};

			try {
				const response = await fetch('/add_product_in_basket', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json',
					},
					body: JSON.stringify(partBasket),
				});

				if (!response.ok) {
					const errorText = await response.text();
					throw new Error(errorText);
				}

				alert("Товар добавлен в корзину!")
			} catch (error) {
				alert(error);

				throw error;
			}
		},
	},
};
</script>
<style>
.page-button {
	background-color: #007bff;
	color: #fff;
	border: none;
	padding: 8px 16px;
	margin: 2px;
	cursor: pointer;
	border-radius: 4px;
}

.active-page {
	background-color: #0056b3;
}

.pagination {
	display: flex;
	justify-content: center;
	align-items: center;
}
</style>