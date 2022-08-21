<template>
  <div class="card">
    <div class="card-header card-header-rose"
    data-header-animation="true">
      <img :src="projection.Movie.Picture" height="350px"
    width="250px"/>
    </div>
    <div class="card-body">
      <h1 class="card-title">{{projection.Movie.Name}}</h1>
      <h3>{{projection.Movie.Plot}}</h3>
      <h3><b>Director:</b>{{projection.Movie.Director}}</h3>
      <h3><b>Actors:</b>{{projection.Movie.Actors}}</h3>
      <div class="card col-md-6">
        <div class="card-header card-header-rose">
          <h3 class="card-title"><b>Projection Information</b></h3>
        </div>
        <div class="card-body">
          <h4><b>{{formatDateTime(projection.Slot)}}</b></h4>
          <h4><b>{{projection.Movie.Duration}}min</b></h4>
          <h4><b>{{projection.CinemaHall.Name}}</b></h4>
          <h4><b>{{projection.Price}} RSD</b></h4>
        </div>
      </div>
      
    </div>
    <div class="card-footer">
    <Button @click="getTickets" class="pull-left">Get Tickets</Button>
    <Button @click="deleteProjection" v-if="user.role === Roles.Worker">Delete</Button>
    </div>
  </div>
</template>

<script>
import Button from '../../../components/Form/Button.vue' 
import moment from 'moment'
import { mapActions } from 'vuex';
import { getUserIdFromToken, getRoleFromToken } from '../../../utils/token';
import {Roles} from '../../../constants.js';

export default {
    components: {
        Button
    },

    data: function () {
      return {
        Roles,
        user: {}
      };
    },
    props: ['projection'],

    methods: {
        ...mapActions({
          delete: "projections/deleteProjection"
        }),

        getTickets(e) {
            e.preventDefault();
            this.$router.push(`/projections/${this.projection.id}`);
        },
        formatDateTime(date) {
            return moment(date).format("DD.MM.YYYY HH:mm");
        },
        deleteProjection() {
          this.delete(this.projection.Id)
        }

    },
    mounted(){
      this.user = {
      id: getUserIdFromToken(),
      role: getRoleFromToken()
      }
    }
}
</script>

<style>

</style>