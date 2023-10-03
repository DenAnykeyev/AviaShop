<script setup>
</script>
<template>
	<h1 class="text-center">Войти</h1>
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