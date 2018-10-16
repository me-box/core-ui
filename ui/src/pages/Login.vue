<template>
	<div class="mdc-card">
		<section class="material-icons mdc-theme--primary-bg">lock</section>
		<form>
			<div class="mdc-text-field" id="url-field">
				<input id="url-field-input"
				       v-model="url"
				       type="text"
				       class="mdc-text-field__input"
				       required
				       autocomplete="username"
				       :autofocus="isMobile">
				<label class="mdc-floating-label" for="url-field-input">Databox URL</label>
				<div class="mdc-line-ripple"></div>
			</div>
			<div class="mdc-text-field" id="password-field">
				<input id="password-field-input"
				       v-model="password"
				       type="password"
				       class="mdc-text-field__input"
				       required
				       autocomplete="current-password"
				       :autofocus="!isMobile">
				<label class="mdc-floating-label" for="password-field-input">Password</label>
				<div class="mdc-line-ripple"></div>
			</div>
			<div style="display: flex">
				<button type="button" v-if="isMobile" class="mdc-button" @click="scan">Scan QR</button>
				<button type="submit" class="mdc-button" @click="login" :disabled="!valid">Login</button>
			</div>
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
				url: "",
				error: ""
			}
		},
		computed: {
			valid: function () {
				return this.url && this.password;
			}
		},
		mounted: function () {
			new MDCTextField(document.querySelector('#url-field'));
			new MDCTextField(document.querySelector('#password-field'));
		},
		created: function() {
			if(!this.isMobile) {
				this.url = window.location.hostname;
			} else {
				this.url = localStorage.getItem('databoxURL');
			}
		},
		methods: {
			scan: function () {
				if(this.isMobile) {
					this.$router.push('/scan');
				}
			},
			login: function () {
				this.$parent.login(this.url, this.password);
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

	.mdc-text-field {
		margin-top: 8px;
	}

	button {
		width: 100%;
		margin: 8px 0;
	}
</style>