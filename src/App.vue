<script setup>
	import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
	<header>
		<nav class="navbar navbar-expand-lg navbar-light bg-light">
			<div class="container-fluid">
				<!-- Логотип и название -->
				<a class="navbar-brand" href="#"> Авиакасса </a>

				<div v-if="isLoggedIn" class="text-center">Привет, {{ username }}</div>

				<!-- Кнопка для развертывания меню на мобильных устройствах -->
				<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav"
					aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
					<span class="navbar-toggler-icon"></span>
				</button>

				<!-- Меню навигации -->
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
						<li class="nav-item">
							<router-link class="nav-link" to="/about">{{this.isLoggedIn}}</router-link>
						</li>
					</ul>

					<!-- Кнопка для регистрации -->
					<button class="btn btn-primary ms-auto">
						<router-link class="nav-link" to="/register"> <i class="bi bi-person-plus"></i> </router-link>
					</button>
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
				isLoggedIn: false
			};
		},
		async created() {
			await this.checkAuth();
		},
		mounted() {
			this.checkAuth();
		},
		methods: {
			async checkAuth() {
				const response = await fetch("/check_auth");

				if (!response.ok) {
                    throw new Error('Не удалось получить данные о авторизации');
                }

				const data = await response.json();

				alert(data);

				if (data.isLoggedIn) {
					this.isLoggedIn = true;
					this.username = data.username;
				} else {
					this.isLoggedIn = false;
					this.username = "";
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