<template>
  <div class="contet">
    <div class="container-fluid">
      <br />
      <br />
      
      <SearchProjectionForm
        @search="onSearchParamsChange($event)"
        :existingParams="searchParams"
      />
      <div v-if="user.role === Roles.Worker" class="row move-button" >
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
import Button from "../../../components/Form/Button.vue";
import toastr from "toastr";
import SearchProjectionForm from "./SearchProjectionForm.vue";
import { getUserIdFromToken, getRoleFromToken } from '../../../utils/token';
import {Roles} from '../../../constants.js';

const initialSearchParams = {
  Name: "",
  Director: "",
  Actor: "",
  Genre: -1,
  Language: "",
  Country: "",
  Duration: 0,
  DateFrom: "",
  DateTo: "",
  Score: 0,
  Year: 0,
};
export default {
  components: {
    ProjectionCard,
    Button,
    SearchProjectionForm,
  },

  data: () => ({
    searchParams: { ...initialSearchParams },
    Roles,
    user: {}
  }),
  computed: {
    ...mapGetters({
      projections: "projections/getProjections",
      result: "projections/getResult",
    }),
  },
  watch: {
    result({ ok, message, label }) {
      if (label !== "delete") return;
      if (ok) {
        toastr.success(message);
        this.fetchProjections();
      } else {
        toastr.error(message);
      }
    },
  },

  methods: {
    ...mapActions({
      fetchProjections: "projections/fetchProjections",
      searchProjections: "projections/searchProjections"
    }),

    onSearchParamsChange(params) {
      this.searchParams = { ...this.searchParams, ...params };
      console.log(this.searchParams)
      if (this.searchParams.DateFrom === "") {
        this.searchParams.DateFrom = "0001-01-01T00:00:00.000Z"
      }

      if(this.searchParams.DateTo === "") {
        this.searchParams.DateTo = "0001-01-01T00:00:00.000Z"
      }
      this.handleSearch();
    },
    handleSearch() {
      this.searchProjections(this.searchParams);
    },
    handleAddProjection() {
      this.$router.push("/projections/add-new-projection");
    },
  },

  mounted() {
    this.fetchProjections();
    this.user = {
      id: getUserIdFromToken(),
      role: getRoleFromToken()
      }
  },
};
</script>

<style>
.move-button {
  margin: 0.2em;
}
</style>
