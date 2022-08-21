<template>
  <div class="card">
    <div class="card-header card-header-rose" data-header-animation="true">
      <img :src="movie.Picture" height="350px" width="250px" />
    </div>
    <div class="card-body">
      <h1 class="card-title">{{ movie.Name }}</h1>
      <h4>{{ movie.Plot }}</h4>
      <h4><b>Director:</b>{{ movie.Director }}</h4>
      <h4><b>Actors:</b>{{ movie.Actors }}</h4>
      <h4><b>Duration:</b>{{ movie.Duration }}min</h4>
      <h4><b>Language:</b>{{movie.Language}}</h4>
      <h4><b>Country:</b>{{movie.Country}}</h4>
      <h4><b>Year:</b>{{movie.Year}}</h4>
      <h4><b>Genre:</b>{{getGenre()}}</h4>    
    </div>
    <div class="card-footer">
      <Button @click="updateMovie" v-if="user.role === Roles.Worker" class="pull-left">Update</Button>
    </div>
  </div>
</template>

<script>
import Button from '../../components/Form/Button.vue' 
import { getAccountIdFromToken, getRoleFromToken } from '../../utils/token';
import {Roles, Genre} from '../../constants.js'
export default {
  components: {
    Button
  },

  data: function () {
    return {
      Roles, 
      Genre,
      user: {}
    };
  }, 
  props: ['movie'],

  methods: {
    getGenre() {
      return Genre[this.movie.Genre] 
    }
  },
  mounted(){
      this.user = {
      id: getAccountIdFromToken(),
      role: getRoleFromToken()
      }
    }
}
</script>

<style>
</style>