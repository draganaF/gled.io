<template>
  <div class="contet">
    <div class="container-fluid">
      <br/>
      <br/>
      <div class="row move-button">
        <Button @click="handleAddMovie">ADD MOVIE</Button>
      </div>
      <div v-for="(movie, index) in movies" :key="index">
        <div class="row">
          <div class="col-md-6 ml-auto mr-auto">
            <MovieCard :movie="movie" />
          </div>
        </div>
      </div>
      
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import MovieCard from "./MovieCard.vue";
import Button from "../../../components/Form/Button.vue";
import toastr from "toastr";

export default {
  components: {
    MovieCard,
    Button,
  },

  computed: {
    ...mapGetters({
      movies: "movies/getMovies",
      result: "movies/getResult",
    }),
  },
  watch: {
    result({ok, message, label}) {
      if (label !== "delete") return;
      if (ok) {
        toastr.success(message);
        this.fetchMovies()
      } else {
        toastr.error(message);
      }
    }
  },

  methods: {
    ...mapActions({
      fetchMovies: "movies/fetchMovies",
    }),
    handleAddMovie() {
      this.$router.push("/movies/add-new-movie");
    },
  },

  mounted() {
    this.fetchMovies();
  },
};
</script>

<style>
.move-button{margin: 1em;}
</style>
