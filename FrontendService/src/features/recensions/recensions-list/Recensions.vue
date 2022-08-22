<template>
  
  <div class="content">

    <RecensionsList
      :movieId=movieId
    />

  </div>

</template>

<script>
import RecensionsList from './RecensionsList.vue';

export default {
  components: { RecensionsList },

  data: () => ({
    movieId: ''
  }),

  mounted() {
    this.movieId = +this.$route.params.id || '';
  }
}
</script>

<style scoped>
</style>