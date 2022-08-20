<template>
  <div class="contet">
    <div class="container-fluid">
      <div v-for="(projection, index) in projections" :key="index">
        <div class="row">
          <div class="col-md-6 ml-uto mr-auto">
            <ProjectionCard :projection="projection"/>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import { mapGetters, mapActions } from 'vuex'
import ProjectionCard from './ProjectionCard.vue'
export default {
  components: {
    ProjectionCard
  }, 

  computed: {
    ...mapGetters({
      projections: 'projections/getProjections',
      result: 'projections/getResult'
    })
  },

  methods: {
    ...mapActions({
      fetchProjections: 'projections/fetchProjections'
    })
  },

  mounted() {
    this.fetchProjections();
  }
}
</script>

<style>
</style>
