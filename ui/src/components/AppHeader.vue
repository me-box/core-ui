<template>
	<div style="display: flex">
		<icon v-bind:name="manifestName" v-bind:displayName="false"/>
		<div>
			<div>{{ manifest.name }}</div>
			<div>{{ manifest.author }}</div>
			<div>{{ manifest.version }}</div>
		</div>
	</div>
	<div>
		<span class="material-icons">star</span>
		<span class="material-icons">star</span>
		<span class="material-icons">star</span>
		<span class="material-icons">star</span>
		<span class="material-icons disabled">star</span>
		&nbsp; Moderate Risk
	</div>
	<div>{{ manifest.description }}</div>
</template>
<script>

	export default {
		name: 'header',
		props: {
			manifest: Object,
			name: String,
			displayName: Boolean,
			updating: Boolean,
			icon: String,
			route: String,
			letter: String
		},
		computed: {
			displayLetter: function () {
				return this.letter || this.name.replace('-', '').replace('_', '').replace("app", '').replace("driver", '').replace('core', '')[0].toUpperCase()
			},
			display: function () {
				return typeof(this.displayName) === "undefined" ? true : this.displayName;
			}
		},
		methods: {
			goto: function () {
				this.$router.push(this.route);
			}
		}
	}

</script>

<style scoped>
	.app-icon {
		width: 128px;
		cursor: pointer;
		user-select: none;
		margin: 8px;
	}

	.app-icon:hover {
		opacity: 0.9;
	}

	.app-icon-letter {
		-webkit-box-align: center;
		-webkit-box-pack: center;
		align-items: center;
		background: #536878;
		border-radius: 2px;
		color: #fff;
		display: -webkit-box;
		display: -ms-flexbox;
		display: flex;
		font-size: 80px;
		height: 128px;
		justify-content: center;
		width: 128px;
		border: 1px #456 solid;
	}

	.text {
		color: #555;
		font-weight: bold;
		width: 128px;
		margin-top: 4px;
		text-align: center;
	}

	@media only screen and (max-width: 768px) {
		.app-icon {
			width: 64px;
		}

		.app-icon-letter {
			font-size: 40px;
			height: 64px;
			width: 64px;
		}

		.text {
			width: 64px;
			font-size: 14px
		}
	}

	.loader {
		color: #ffffff;
		font-size: 70px;
		text-indent: -9999em;
		overflow: hidden;
		width: 1em;
		height: 1em;
		border-radius: 50%;
		margin: 72px auto;
		position: relative;
		-webkit-transform: translateZ(0);
		-ms-transform: translateZ(0);
		transform: translateZ(0);
		-webkit-animation: load6 1.7s infinite ease, round 1.7s infinite ease;
		animation: load6 1.7s infinite ease, round 1.7s infinite ease;
	}

	@-webkit-keyframes load6 {
		0% {
			box-shadow: 0 -0.83em 0 -0.4em, 0 -0.83em 0 -0.42em, 0 -0.83em 0 -0.44em, 0 -0.83em 0 -0.46em, 0 -0.83em 0 -0.477em;
		}
		5%,
		95% {
			box-shadow: 0 -0.83em 0 -0.4em, 0 -0.83em 0 -0.42em, 0 -0.83em 0 -0.44em, 0 -0.83em 0 -0.46em, 0 -0.83em 0 -0.477em;
		}
		10%,
		59% {
			box-shadow: 0 -0.83em 0 -0.4em, -0.087em -0.825em 0 -0.42em, -0.173em -0.812em 0 -0.44em, -0.256em -0.789em 0 -0.46em, -0.297em -0.775em 0 -0.477em;
		}
		20% {
			box-shadow: 0 -0.83em 0 -0.4em, -0.338em -0.758em 0 -0.42em, -0.555em -0.617em 0 -0.44em, -0.671em -0.488em 0 -0.46em, -0.749em -0.34em 0 -0.477em;
		}
		38% {
			box-shadow: 0 -0.83em 0 -0.4em, -0.377em -0.74em 0 -0.42em, -0.645em -0.522em 0 -0.44em, -0.775em -0.297em 0 -0.46em, -0.82em -0.09em 0 -0.477em;
		}
		100% {
			box-shadow: 0 -0.83em 0 -0.4em, 0 -0.83em 0 -0.42em, 0 -0.83em 0 -0.44em, 0 -0.83em 0 -0.46em, 0 -0.83em 0 -0.477em;
		}
	}

	@keyframes load6 {
		0% {
			box-shadow: 0 -0.83em 0 -0.4em, 0 -0.83em 0 -0.42em, 0 -0.83em 0 -0.44em, 0 -0.83em 0 -0.46em, 0 -0.83em 0 -0.477em;
		}
		5%,
		95% {
			box-shadow: 0 -0.83em 0 -0.4em, 0 -0.83em 0 -0.42em, 0 -0.83em 0 -0.44em, 0 -0.83em 0 -0.46em, 0 -0.83em 0 -0.477em;
		}
		10%,
		59% {
			box-shadow: 0 -0.83em 0 -0.4em, -0.087em -0.825em 0 -0.42em, -0.173em -0.812em 0 -0.44em, -0.256em -0.789em 0 -0.46em, -0.297em -0.775em 0 -0.477em;
		}
		20% {
			box-shadow: 0 -0.83em 0 -0.4em, -0.338em -0.758em 0 -0.42em, -0.555em -0.617em 0 -0.44em, -0.671em -0.488em 0 -0.46em, -0.749em -0.34em 0 -0.477em;
		}
		38% {
			box-shadow: 0 -0.83em 0 -0.4em, -0.377em -0.74em 0 -0.42em, -0.645em -0.522em 0 -0.44em, -0.775em -0.297em 0 -0.46em, -0.82em -0.09em 0 -0.477em;
		}
		100% {
			box-shadow: 0 -0.83em 0 -0.4em, 0 -0.83em 0 -0.42em, 0 -0.83em 0 -0.44em, 0 -0.83em 0 -0.46em, 0 -0.83em 0 -0.477em;
		}
	}

	@-webkit-keyframes round {
		0% {
			-webkit-transform: rotate(0deg);
			transform: rotate(0deg);
		}
		100% {
			-webkit-transform: rotate(360deg);
			transform: rotate(360deg);
		}
	}

	@keyframes round {
		0% {
			-webkit-transform: rotate(0deg);
			transform: rotate(0deg);
		}
		100% {
			-webkit-transform: rotate(360deg);
			transform: rotate(360deg);
		}
	}
</style>
