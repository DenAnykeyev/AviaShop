import { createRouter, createWebHistory } from 'vue-router'
import ProductsView from '../views/Products.vue'
import AboutView from '../views/About.vue'
import HomeView from '../views/Home.vue'
import ProductsEditorView from '../views/ProductsEditor.vue'
import RegisterView from '../views/Register.vue'
import LoginView from '../views/Login.vue'
import BasketView from '../views/Basket.vue'

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			name: 'home',
			component: HomeView
		},
		{
			path: '/about',
			name: 'about',
			component: AboutView
		},
		{
			path: '/products',
			name: 'products',
			component: ProductsView
		},
		{
			path: '/products_edit',
			name: 'products_edit',
			component: ProductsEditorView
		},
		{
			path: '/register',
			name: 'register',
			component: RegisterView
		},
		{
			path: '/login',
			name: 'login',
			component: LoginView
		},
		{
			path: '/basket',
			name: 'basket',
			component: BasketView
		}
	]
})

export default router