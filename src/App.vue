<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
	<header>
		<nav class="navbar navbar-expand-lg navbar-light bg-light">
			<div class="container-fluid">
				<a class="navbar-brand" href="#"> Авиакасса </a>

				<div v-if="isLoggedIn" class="text-center">Привет, {{ username }}</div>

				<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav"
					aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
					<span class="navbar-toggler-icon"></span>
				</button>

				<div class="collapse navbar-collapse" id="navbarNav">
					<ul class="navbar-nav">
						<li class="nav-item">
							<router-link class="nav-link" to="/"><i class="bi bi-house"></i> Главная</router-link>
						</li>
						<li class="nav-item">
							<router-link class="nav-link" to="/products"><i class="bi bi-star"></i> Товары</router-link>
						</li>
						<li class="nav-item">
							<router-link class="nav-link" to="/about"><i class="bi bi-person-circle"></i> О
								нас</router-link>
						</li>
						<li v-if="this.isAdmin === true" class="nav-item">
							<router-link class="nav-link" to="/products_edit">Редактор товаров</router-link>
						</li>
						<li class="nav-item">
							<router-link class="nav-link" to="/basket">Корзина</router-link>
						</li>
					</ul>

					<div v-if="isLoggedIn === false" class="d-flex align-items-center ms-auto">
						<button class="btn btn-primary">
							<router-link class="nav-link" to="/register"> <i class="bi bi-person-plus"></i> Войти
							</router-link>
						</button>
					</div>

					<button v-else class="btn btn-primary ms-auto bi bi-door-closed" @click="logout"> {{ name }}</button>
				</div>
			</div>
		</nav>
	</header>
	<main class="container">
		<RouterView />
	</main>
	<footer>
	</footer>
</template>
<script>
export default {
	data() {
		return {
			isLoggedIn: false,
			name: "",
			rules: "",
			isAdmin: false
		};
	},
	mounted() {
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
				if (this.rules == "admin") {
					this.isAdmin = true;
				}
			} else {
				this.isLoggedIn = false;
				this.name = "";
			}
		},

		async logout() {
			const confirmDelete = confirm('Вы уверены, что хотите выйти из аккаунта?');
			if (confirmDelete) {
				try {
					const response = await fetch("/logout_user", {
						method: "POST",
					});

					if (response.ok) {
						this.isLoggedIn = false;
						this.name = "";

						window.location.reload();
					} else {
						alert("Ошибка выхода из аккаунта");
					}

				} catch (error) {
					alert("Ошибка выхода из аккаунта", error);
				}
			}
		},
	},
};
</script>
<style scoped>
header {
	background-color: #343a40;
	color: white;
}

.navbar-brand {
	font-size: 1.5rem;
}

.nav-link {
	color: #000000 !important;
	font-size: 1.2rem;
}

.btn-primary {
	background-color: #007bff;
	border: none;
}

.btn-primary:hover {
	background-color: #0056b3;
}
</style>