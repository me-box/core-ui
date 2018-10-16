<template>
	<div>
		<div style="display: flex; margin-bottom: 8px">
			<icon v-bind:name="datasource" v-bind:displayName="false"/>
			<div style="margin-left: 8px">
				<div>{{ datasource }}</div>
			</div>
		</div>
		<div style="display: flex; justify-content: space-between;">
			<div class="mdc-button icon-button">
				<div class="material-icons">share</div>
				Share
			</div>
			<div class="mdc-button icon-button">
				<div class="material-icons">cloud_download</div>
				Export
			</div>
			<div class="mdc-button icon-button">
				<div class="material-icons">check_box</div>
				Enabled
			</div>
			<div class="mdc-button icon-button" style="color: firebrick" @click="deletePrompt">
				<div class="material-icons">delete</div>
				Delete
			</div>
		</div>
		<aside id="delete-dialog"
		       class="mdc-dialog"
		       role="alertdialog"
		       aria-modal="true"
		       aria-describedby="delete-dialog-content">
			<div class="mdc-dialog__container">
				<div class="mdc-dialog__surface">
					<section id="delete-dialog-content" class="mdc-dialog__content">
						{{ 'Permanently delete ' + this.datasource + ' data?' }}
					</section>
					<footer class="mdc-dialog__actions">
						<button class="mdc-button mdc-dialog__footer__button mdc-dialog__footer__button--cancel" @click="onDecline">Cancel</button>
						<button class="mdc-button mdc-dialog__footer__button mdc-dialog__footer__button--accept mdc-dialog__action" style="color: firebrick" @click="deleteData">Delete</button>
					</footer>
				</div>
			</div>
			<div class="mdc-dialog__scrim"></div>
		</aside>
	</div>
</template>
<script>
	import {MDCDialog} from '@material/dialog';
	import Icon from '../components/AppIcon.vue'

	export default {
		name: 'databoxStatus',
		props: ['datasource'],
		components: {
			Icon,
		},
		data() {
			//get data from api later
			return {
				dataSources: [],
				timerID: 0,
				deleteDialog: null,
			}
		},
		mounted() {
			this.$parent.setTitle(this.datasource + ' Data');
			this.deleteDialog = new MDCDialog(document.querySelector('#delete-dialog'));
		},
		destroyed: function () {
			clearInterval(this.timerID)
		},
		methods: {
			deletePrompt() {
				this.deleteDialog.open();
			}			,
			deleteData() {
				console.log('accepted');
			},
			onDecline() {
			},
		}
	}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
	.icon-button {
		flex-direction: column;
		height: auto;
		padding: 8px;
	}

	.icon-button > .material-icons {
		padding-bottom: 8px;
	}
</style>
