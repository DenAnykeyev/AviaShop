<script setup>
</script>

<template>
	<div class="container">
		<div id="carouselExampleCaptions" class="carousel slide" data-bs-ride="carousel">
			<div class="carousel-indicators">
				<button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="0" class="active"
					aria-current="true" aria-label="Slide 1"></button>
				<button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="1"
					aria-label="Slide 2"></button>
				<button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="2"
					aria-label="Slide 3"></button>
			</div>
			<div class="carousel-inner">
				<div class="carousel-item active">
					<img src="assets/img/карусель1.jpg" class="d-block w-100 carousel-image-round" alt="...">
					<div class="carousel-caption d-none d-md-block">
						<h5>Наши самолеты</h5>
						<p>Хорошего полета!</p>
					</div>
				</div>
				<div class="carousel-item">
					<img src="assets/img/карусель2.jpg" class="d-block w-100 carousel-image-round" alt="...">
					<div class="carousel-caption d-none d-md-block">
						<h5>Супер крутой самолет</h5>
						<p>Нереально крутого полета!</p>
					</div>
				</div>
				<div class="carousel-item">
					<img src="assets/img/карусель3.jpg" class="d-block w-100 carousel-image-round" alt="...">
					<div class="carousel-caption d-none d-md-block">
						<h5>Самолет премиум класса</h5>
						<p>Самолет для самых крутых!</p>
					</div>
				</div>
			</div>
			<button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleCaptions"
				data-bs-slide="prev">
				<span class="carousel-control-prev-icon" aria-hidden="true"></span>
				<span class="visually-hidden">Назад</span>
			</button>
			<button class="carousel-control-next" type="button" data-bs-target="#carouselExampleCaptions"
				data-bs-slide="next">
				<span class="carousel-control-next-icon" aria-hidden="true"></span>
				<span class="visually-hidden">Вперед</span>
			</button>
		</div>
		<div class="my-box">
			<div class="text-center">
				<h1 class="text-center">Авиакасса</h1>
				<p class="text-center">Добро пожаловать!<br>Мы занимаемся тем, что продаем авиабилеты на любые рейсы...
				</p>
				<img src="assets/img/main.png" class="img-fluid rounded">
				<p><br></p>
			</div>
		</div>

		<div class="my-box">
			<h1 class="text-center">Наши отзывы</h1>
			<div class="review-box">
				<img src="assets/img/user.jpg" alt="Пользователь" class="review-user-image">
				<span class="review-text">Спасибо Вам за существование!</span>
			</div>

			<div class="review-box">
				<img src="assets/img/user.jpg" alt="Пользователь" class="review-user-image">
				<span class="review-text">Замечательный сервис!</span>
			</div>

			<div class="review-box">
				<img src="assets/img/user.jpg" alt="Пользователь" class="review-user-image">
				<span class="review-text">Слишком долгая загрузка...</span>
			</div>
		</div>

	</div>
</template>
<script>
</script>
<style>
	.carousel-image-round {
		border-radius: 50%;
	}

	.my-box {
		background-color: #f1f1f1;
		border-radius: 10px;
		padding: 20px;
		margin-top: 20px;
	}

	.review-box {
		display: flex;
		align-items: center;
		justify-content: flex-start;
		background-color: #bbb9b9;
		border-radius: 10px;
		padding: 20px;
		margin-top: 10px;
		margin-bottom: 10px;

	}

	.review-user-image {
		width: 80px;
		border-radius: 5px;
		box-shadow: 0 0 0 3px #0780ff, 0 0 13px #333;
	}

	.review-text {
		display: block;
		text-align: left;
		margin-left: 20px;
		font-size: 16px;
	}
</style>