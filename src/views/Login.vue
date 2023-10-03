<script setup>
</script>
<template>
	<h1 class="text-center">Войти</h1>
	<form @submit.prevent="loginUser" class="my-form">
		<div class="form-group">
			<label for="username">Имя:</label>
			<input type="text" id="username" v-model="username" class="form-control" required />
		</div>

		<div class="form-group">
			<label for="password">Пароль:</label>
			<input type="password" id="password" v-model="password" class="form-control" required />
		</div>
		<div class="text-center"><button type="submit" class="btn btn-primary">Отправить</button></div>
		
	</form>
</template>
<script>
	export default {
		data() {
			return {
				isLoggedIn: false,
			};
		},
		async mounted() {
			try {
				await this.checkAuth();

				if (this.isLoggedIn) {
					window.location.href = '/';
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
				} else {
					this.isLoggedIn = false;
					this.name = "";
				}
			},
			async loginUser() {
				const username = this.username;
				const password = this.password;

				const user = {
					name: username,
					password: password
				};

				try {
					const response = await fetch('/login_user', {
						method: 'POST',
						headers: {
							'Content-Type': 'application/json',
						},
						body: JSON.stringify(user),
					});

					if (!response.ok) {
						const errorText = await response.text();
						throw new Error(errorText);
                    }

					window.location.href = '/';
				} catch (error) {
					alert(error);

					throw error;
				}

				
			},
		},
	};
</script>
<style>
	.my-form {
		max-width: 400px;
		margin: 0 auto;
		padding: 20px;
		border: 1px solid #ccc;
		border-radius: 5px;
		background-color: #f9f9f9;
	}

	.form-group {
		margin-bottom: 20px;
	}
</style>