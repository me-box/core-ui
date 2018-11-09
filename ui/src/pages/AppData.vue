<template>
	<div style="max-width: 100%">
		<div style="display: flex; margin-bottom: 8px">
			<icon v-bind:name="app" v-bind:displayName="false"/>
			<div style="margin-left: 8px">
				<div>{{ app }}</div>
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
		<div class="mdc-list mdc-list--two-line" data-mdc-auto-init="MDCList">
			<div class="mdc-list-item" v-for="source in dataSources" v-on:click="goto(source.dsType)">
				<div class="mdc-list-item__graphic material-icons" aria-hidden="true">folder</div>
				<div class="mdc-list-item__text">
                    <div class="mdc-list-item__primary-text">{{source.description}}</div>
                    <div class="mdc-list-item__secondary-text">{{source.dsType}}</div>
                </div>
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
	import testdata from '../testData/DataSources.json'
	import Icon from '../components/AppIcon.vue'

	export default {
		name: 'databoxStatus',
		props: ['app'],
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
			this.$parent.setTitle(this.$route.params.app + ' Data');
			this.deleteDialog = new MDCDialog(document.querySelector('#delete-dialog'));
			this.loadData();
			this.timerID = setInterval(() => {
				this.loadData();
			}, 5000);

		},
		destroyed: function () {
			clearInterval(this.timerID)
		},
		methods: {
			deleteData() {
				console.log('accepted');
				this.deleteDialog.close();
			},
			onDecline() {
				this.deleteDialog.close();
			},

			deletePrompt() {
				this.deleteDialog.open();
			},

			loadData: function () {
				this.$parent.apiRequest('/core-ui/ui/api/dataSources', testdata)
					.then(json => {
						//this.dataSources = json;
						let getValFromHypercat = (hypercat, match) => {
							return hypercat["item-metadata"].filter((itm) => {
								return itm.rel === match
							})[0].val;
						};

						let filter = "://" + this.app + "-core-store";
						let sources = [];
						json.forEach((ds) => {
							if (ds.href.indexOf(filter) !== -1) {
								sources.push({
									name: getValFromHypercat(ds, "urn:X-databox:rels:hasDatasourceid"),
									description: getValFromHypercat(ds, "urn:X-hypercat:rels:hasDescription:en"),
									vendor: getValFromHypercat(ds, "urn:X-databox:rels:hasVendor"),
									dsType: getValFromHypercat(ds, "urn:X-databox:rels:hasType"),
									href: ds["href"],
									hypercat: ds
								})
							}
						});
						this.dataSources = sources;
					})
			},
			goto: function (sourceName) {
				this.$router.push('/datasource/' + sourceName);
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
