<template>
	<div class="app-icon" :class="{ 'hover': route && !updating }" @click="goto()">
		<div class="app-icon-letter" :class="{ 'material-icons': icon }">
			<svg v-if="updating" class="spinner" viewBox="0 0 66 66" xmlns="http://www.w3.org/2000/svg">
				<circle class="path" fill="none" stroke-width="6" stroke-linecap="round" cx="33" cy="33" r="30">
				</circle>
			</svg>
			<template v-else-if="icon">{{ icon }}</template>
			<template v-else>{{ displayLetter }}</template>
			<div v-if="banner" class="banner">{{ banner }}</div>
		</div>
		<div v-if="display" class="text">{{ name }}</div>
	</div>
</template>
<script>
	export default {
		name: 'icon',
		props: {
			name: String,
			displayName: Boolean,
			updating: Boolean,
			icon: String,
			route: String,
			letter: String,
			banner: String
		},
		computed: {
			displayLetter() {
				return this.letter || this.name.replace('-', '').replace('_', '').replace("app", '').replace("driver", '').replace('core', '')[0].toUpperCase()
			},
			display() {
				return typeof(this.displayName) === "undefined" ? true : this.displayName;
			}
		},
		methods: {
			goto() {
				if (this.route && !this.updating) {
					this.$router.push(this.route);
				}
			}
		}
	}

</script>

<style scoped lang="scss">
	.app-icon {
		width: 128px;
		user-select: none;
	}

	.hover {
		cursor: pointer;
	}

	.hover:hover {
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
		position: relative;
	}

	.banner {
		position: absolute;
		bottom: 0;
		left: 0;
		right: 0;
		font-size: 12px;
		text-transform: uppercase;
		background: rgba(0,0,0,0.2);
		text-align: center;
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

		.spinner {
			width: 30px;
			height: 30px;
		}
	}

	// Here is where the magic happens

	$offset: 187;
	$duration: 1.4s;

	.spinner {
		animation: rotator $duration linear infinite;
		width: 65px;
		height: 65px;
	}

	@keyframes rotator {
		0% {
			transform: rotate(0deg);
		}
		100% {
			transform: rotate(270deg);
		}
	}

	.path {
		stroke-dasharray: $offset;
		stroke-dashoffset: 0;
		transform-origin: center;
		animation: dash $duration ease-in-out infinite,
		colors ($duration*4) ease-in-out infinite;
	}

	@keyframes colors {
		0% {
			stroke: #FFF;
		}
	}

	@keyframes dash {
		0% {
			stroke-dashoffset: $offset;
		}
		50% {
			stroke-dashoffset: $offset/4;
			transform: rotate(135deg);
		}
		100% {
			stroke-dashoffset: $offset;
			transform: rotate(450deg);
		}
	}
</style>
