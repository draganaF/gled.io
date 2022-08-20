<template>
  <div class="card">
    <div class="card-header card=header-rose"
    data-header-animation="true">
      <img :src="projection.Movie.Picture"/>
    </div>
    <div class="card-body">
      <h4 class="card-title">{{projection.Movie.Name}}</h4>
      <p>Plot:{{projection.Movie.Plot}}</p>
      <p>Cinema Hall: {{projection.CinemaHall.Name}}</p>
      <p>Time:{{projection.Slot}}</p>
    </div>
    <div class="card-footer">
    <Button @click="getTickets" class="pull-right">Get Tickets</Button>
    </div>
  </div>
</template>

<script>
import Button from '../../components/Form/Button.vue' 
export default {
    components: {
        Button
    },

    props: ['projection'],

    methods: {

        getTickets(e) {
            e.preventDefault();
            this.$router.push(`/projections/${this.projection.id}`);
        }
    }
}
</script>

<style>

</style>