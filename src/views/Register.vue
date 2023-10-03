<script setup>
</script>
<template>
	<h1 class="text-center">Зарегистрироваться</h1>
	<form @submit.prevent="registerUser" class="my-form">
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

			};
		},
		mounted() {

		},
		methods: {

			// Меняй ТОЛЬКО ТУТ КОД, В СЕТАП НЕ НАДО
			async registerUser() {
				// Получите имя пользователя и пароль из полей ввода
				const username = this.username;
				const password = this.password;

				// Создайте объект с данными нового пользователя
				const newUser = {
					name: username,
					password: password,
					rules: "user",
				};

				try {
					// Отправьте POST-запрос на сервер с данными пользователя
					const response = await fetch('/register_user', {
						method: 'POST',
						headers: {
							'Content-Type': 'application/json',
						},
						body: JSON.stringify(newUser),
					});

					if (!response.ok) {
                        // Получите текст ошибки из ответа бэкенда
						const errorText = await response.text();
						throw new Error(errorText);
                    }

					alert("Вы были успешно зарегистрированы")
					
					return true;
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