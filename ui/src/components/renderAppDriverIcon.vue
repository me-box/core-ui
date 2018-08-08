<template>
  <div>
  <div class="app-icon-letter">
      {{ letter }}
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
  },
  computed: {
      letter: function () {
          return this.name.replace('-','').replace("app",'').replace("driver",'').replace('core','')[0].toUpperCase()
      },
      display: function () {
          return typeof(this.displayName) == "undefined" ? true : this.displayName;
      }
  }
}

</script>

<!--  "scoped" CSS -->
<style scoped>

.app-icon-letter {
	display: flex;
	justify-content: center;
	align-items: center;
	color: #fff;
	border-radius: 5px;
	font-size: 98px;
	background: #3f51b5;
	background: linear-gradient(#6f7dcd, #3f51b5);
	width: 134px;
	height: 134px;
    cursor: pointer;
}

.text {
    text-decoration: underline;
    font-weight: bold;
    width: 100%;
    text-align: center;
}
</style>
