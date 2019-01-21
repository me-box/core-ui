<template>
	<div class="mdc-card">
		<section class="material-icons mdc-theme--primary-bg"
		         style="border-top-left-radius: 4px; border-top-right-radius: 4px">lock
		</section>
		<form @submit.prevent="login">
			<div class="mdc-text-field" id="url-field">
				<input id="url-field-input"
				       v-model="url"
				       type="text"
				       class="mdc-text-field__input"
				       required
				       autocomplete="username"
				       :disabled="connecting"
				       :autofocus="isMobile">
				<label class="mdc-floating-label" :class="{ 'mdc-floating-label--float-above' : url }"
				       for="url-field-input">Databox URL</label>
				<div class="mdc-line-ripple"></div>
			</div>
			<div v-if="urlError" class="mdc-text-field-helper-text--persistent mdc-text-field-helper-text error"
			     aria-hidden="true">
				{{ urlError }}
			</div>
			<div class="mdc-text-field" id="password-field">
				<input id="password-field-input"
				       v-model="password"
				       type="password"
				       class="mdc-text-field__input"
				       required
				       autocomplete="current-password"
				       :disabled="connecting"
				       :autofocus="!isMobile">
				<label class="mdc-floating-label" for="password-field-input">Password</label>
				<div class="mdc-line-ripple"></div>
			</div>
			<div v-if="passwordError" class="mdc-text-field-helper-text--persistent mdc-text-field-helper-text error"
			     aria-hidden="true">
				{{ passwordError }}
			</div>

			<div v-if="connecting" style="display: flex; justify-content: center; padding: 16px">
				<Spinner/>
			</div>
			<div v-else style="display: flex">
				<button type="button" v-if="isMobile" class="mdc-button" @click="scan">Scan QR</button>
				<button type="submit" class="mdc-button" :disabled="!valid">Login</button>
			</div>
		</form>

		<section class="mdc-theme--primary-bg"></section>
	</div>
</template>
<script>
	import {MDCTextField} from '@material/textfield';
	import Spinner from '../components/Spinner.vue';

	export default {
		name: 'logIn',
		props: {
			url: {
				default: function () {
					return this.$parent.databoxUrl;
				},
				type: String
			},
			ipList: {
				default: [],
				type: Array
			},
			password: {
				default: null,
				type: String
			},
			autoLogin: {
				default: false,
				type: Boolean
			}
		},
		components: {Spinner},
		data() {
			return {
				passwordError: "",
				urlError: "",
				connecting: false
			}
		},
		computed: {
			valid() {
				return this.url && this.password;
			}
		},
		mounted() {
			this.$parent.title = null;
			this.$parent.backRoute = null;

			new MDCTextField(document.querySelector('#url-field'));
			new MDCTextField(document.querySelector('#password-field'));

			if (this.autoLogin) {
				this.loginList();
			}
		},
		methods: {
			scan() {
				if (this.isMobile) {
					this.$router.push('/scan');
				}
			},
			async loginList() {
				this.connecting = true;
				for (const ip of this.ipList) {
					try {
						this.url = ip;
						await this.$parent.login(ip, this.password);
						return;
					} catch (error) {
						console.log(error);
					}
				}
				this.connecting = false;
				this.urlError = "";
				this.passwordError = "Login Failed.";
			},
			login() {
				this.connecting = true;
				return this.$parent.login(this.url, this.password)
					.catch((error) => {
						console.log(error);
						this.connecting = false;
						this.passwordError = "";
						this.urlError = "";
						if (error instanceof TypeError) {
							this.urlError = "Could not find a Databox here";
						} else if (error.status === 401) {
							this.passwordError = "Login Failed. Check password.";
						} else if (error.status === 404) {
							this.urlError = "Could not find a Databox here";
						}
					})
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

	.error {
		color: #B33 !important;
	}

	.mdc-text-field {
		margin-top: 8px;
	}

	button {
		width: 100%;
		margin: 8px 0;
	}
</style>