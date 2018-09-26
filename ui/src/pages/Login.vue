<template>
	<div class="mdc-card">
		<section class="material-icons mdc-theme--primary-bg">lock</section>
		<form class="form">
			<div class="mdc-text-field" id="url-field">
				<input type="text" id="url-field-input" class="mdc-text-field__input" required v-model="url">
				<label class="mdc-floating-label" for="url-field-input">Databox URL</label>
				<div class="mdc-line-ripple"></div>
			</div>
			<div class="mdc-text-field" id="password-field">
				<input type="password" id="password-field-input" class="mdc-text-field__input" required v-model="password">
				<label class="mdc-floating-label" for="password-field-input">Password</label>
				<div class="mdc-line-ripple"></div>
			</div>
			<button class="mdc-button" @click="login">Login</button>
		</form>

		<section class="mdc-theme--primary-bg"></section>
	</div>
</template>
<script>
	import {MDCTextField} from '@material/textfield';

	export default {
		name: 'logIn',
		props: {},
		data: function () {
			return {
				password: "",
				url: "127.0.0.1",
				error: ""
			}
		},
		mounted: function () {
			new MDCTextField(document.querySelector('#url-field'));
			new MDCTextField(document.querySelector('#password-field'));
		},
		methods: {
			login: function () {
				//let _this = this;
				fetch('https://' + this.url + '/api/connect', {
					method: "GET",
					credentials: "include",
					mode: "cors",
					headers: {
						'Authorization': "Token " + this.password,
					},
				})
					.then((response) => {
						return response.text()
					})
					.then((body) => {
						console.log(body);
						if (body === "connected") {
							localStorage.setItem('databoxAuthenticated', "true");
							this.$parent.authenticated = true;
							this.$router.push("/")
						} else {
							return Promise.reject("Auth failed")
						}
					})
					.catch((error) => {
						console.log(error);
						this.error = error.toString();
						let devmode = localStorage.getItem('dev');
						if (devmode === "true") {
							localStorage.setItem('databoxURL', this.url);
							localStorage.setItem('databoxAuthenticated', "true");
							this.$parent.authenticated = true;
							this.$router.push("/")
						} else {
							this.$parent.authenticated = false;
							localStorage.setItem('databoxAuthenticated', "false");
						}
					});
			}
		}
	}
</script>
<style scoped>
	section {
		color: white;
		display: block;
		padding: 8px;
		text-align: center;
	}

	form {
		padding: 4px 32px;
		display: flex;
		flex-direction: column;
	}

	button {
		width: 100%;
		margin: 8px 0;
	}
</style>