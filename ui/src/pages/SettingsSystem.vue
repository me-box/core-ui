<template>
	<div id="settingList">
		<div class="mdc-text-field" id="databox-name-field">
			<input type="text" id="my-text-field" class="mdc-text-field__input" required v-model="name">
			<label class="mdc-floating-label" for="my-text-field">Databox Name</label>
			<div class="mdc-line-ripple"></div>
		</div>
		<aside id="password-dialog"
		       class="mdc-dialog"
		       role="alertdialog"
		       aria-modal="true"
		       aria-labelledby="rename-dialog-title"
		       aria-describedby="rename-dialog-content">
			<div class="mdc-dialog__container">
				<div class="mdc-dialog__surface">
					<header id="rename-dialog-title" class="mdc-dialog__title">Set New Password</header>
					<section id="rename-dialog-content" class="mdc-dialog__content">
						<div class="mdc-text-field" id="password-field" style="width: 100%; margin-bottom: 8px">
							<input type="password" id="password-input" class="mdc-text-field__input" required
							       v-model="password">
							<label class="mdc-floating-label" for="password-input">Password</label>
							<div class="mdc-line-ripple"></div>
						</div>
						<div class="mdc-text-field" id="password-confirm-field" style="width: 100%">
							<input type="password" id="password-confirm-input" class="mdc-text-field__input" required
							       v-model="passwordConfirm">
							<label class="mdc-floating-label" for="password-confirm-input">Confirm Password</label>
							<div class="mdc-line-ripple"></div>
						</div>
					</section>
					<footer class="mdc-dialog__actions">
						<button class="mdc-button mdc-dialog__footer__button mdc-dialog__footer__button--cancel"
						        @click="closePasswordDialog">Cancel
						</button>
						<button class="mdc-button mdc-dialog__footer__button mdc-dialog__footer__button--accept mdc-dialog__action"
						        @click="closePasswordDialog">Change Password
						</button>
					</footer>
				</div>
			</div>
			<div class="mdc-dialog__scrim"></div>
		</aside>
		<button class="mdc-button" @click="openPasswordDialog">Change Password</button>
		<button class="mdc-button" @click="logout">Log Out</button>
		
		<div><img alt="QR Code for Mobile Login" src="/core-ui/ui/api/qrcode.png"></div>
	</div>
</template>
<script>
	import {MDCDialog} from '@material/dialog';
	import {MDCTextField} from '@material/textfield';

	export default {
		name: 'settings',
		props: {},
		data: function () {
			return {
				name: "Databox",
				passwordDialog: null,
				password: "",
				passwordConfirm: "",
				error: ""
			}
		},
		mounted: function () {
			this.$parent.title = "Databox Settings";
			this.$parent.backRoute = "/settings";
			new MDCTextField(document.querySelector('#databox-name-field'));
			new MDCTextField(document.querySelector('#password-field'));

			new MDCTextField(document.querySelector('#password-confirm-field'));
			this.passwordDialog = new MDCDialog(document.querySelector('#password-dialog'));
		},
		methods: {
			closePasswordDialog: function () {
				this.passwordDialog.close();
			},
			logout: function () {
				this.$parent.logout();
			},
			openPasswordDialog: function () {
				//this.renameDialog.lastFocusedTarget = evt.target;
				this.passwordDialog.open();
			},

			goto: function (path) {
				this.$router.push(path);
			}
		}
	}
</script>

<style scoped>
	#settingList {
		display: flex;
		flex-direction: column;
		margin: 24px;
	}

	@media (max-width: 479px) {
		#settingList {
			margin: 16px;
		}
	}
</style>