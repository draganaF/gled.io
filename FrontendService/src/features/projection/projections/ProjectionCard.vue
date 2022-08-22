<template>
  <div class="card">
    <div class="card-header card-header-rose"
    data-header-animation="true">
      <img :src="projection.Movie.Picture" height="350px"
    width="250px"/>
    </div>
    <div class="card-body">
      <h1 class="card-title">{{projection.Movie.Name}}</h1>
      <Stars :numOfStars="score"/>
      <h4>{{projection.Movie.Plot}}</h4>
      <h4><b>Director:</b>{{projection.Movie.Director}}</h4>
      <h4><b>Actors:</b>{{projection.Movie.Actors}}</h4>
      <h4><b>Language:</b>{{projection.Movie.Language}}</h4>
      <h4><b>Country:</b>{{projection.Movie.Country}}</h4>
      <h4><b>Year:</b>{{projection.Movie.Year}}</h4>
      <h4><b>Genre:</b>{{getGenre()}}</h4>
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
    <Button @click="getTickets" v-if="!!user.role" class="pull-left">Get Tickets</Button>
    <Button @click="deleteProjection" v-if="user.role === 'Worker'">Delete</Button>
    </div>
  </div>
</template>

<script>
import Button from '../../../components/Form/Button.vue' 
import moment from 'moment'
import { mapActions, mapGetters } from 'vuex';
import { getUserIdFromToken, getRoleFromToken } from '../../../utils/token';
import {Roles, Genre} from '../../../constants.js';
import Stars from '../../../components/Rating/Stars.vue';


export default {
    components: {
    Button,
    Stars
},

    data: function () {
      return {
        Roles,
        user: {}
      };
    },
    props: ['projection'],
    computed: {
        ...mapGetters({
        score: "recensions/getScore",
      })
    },
    methods: {
        ...mapActions({
          delete: "projections/deleteProjection",
          fetchScore: "recensions/fetchScoreByMovieId"
        }),
        getGenre() {
          return Genre[this.projection.Movie.Genre] 
        },
        getTickets(e) {
            e.preventDefault();
            this.$router.push(`/projections/tickets/${this.projection.Id}`);
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
      
      this.fetchScore(this.projection.Movie.Id)
    }
}
</script>

<style>

</style>