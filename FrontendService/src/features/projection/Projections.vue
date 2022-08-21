<template>
  <div class="contet">
    <div class="container-fluid">
      <br/>
      <br/>
      <div class="row move-button">
        <Button @click="handleAddProjection">Add projection</Button>
      </div>
      <div v-for="(projection, index) in projections" :key="index">
        <div class="row">
          <div class="col-md-6 ml-auto mr-auto">
            <ProjectionCard :projection="projection" />
          </div>
        </div>
      </div>
      
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import ProjectionCard from "./ProjectionCard.vue";
import Button from "../../components/Form/Button.vue";
import toastr from "toastr";

export default {
  components: {
    ProjectionCard,
    Button,
  },

  computed: {
    ...mapGetters({
      projections: "projections/getProjections",
      result: "projections/getResult",
    }),
  },
  watch: {
    result({ok, message, label}) {
      if (label !== "delete") return;
      if (ok) {
        toastr.success(message);
        this.fetchProjections()
      } else {
        toastr.error(message);
      }
    }
  },

  methods: {
    ...mapActions({
      fetchProjections: "projections/fetchProjections",
    }),
    handleAddProjection() {
      this.$router.push("/projections/add-new-projection");
    },
  },

  mounted() {
    this.fetchProjections();
  },
};
</script>

<style>
.move-button{margin: 1em;}
</style>
