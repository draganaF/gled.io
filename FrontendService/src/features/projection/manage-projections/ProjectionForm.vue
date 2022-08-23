<template>
  <div class="content">
    <div class="container-fluid">
      <div class="row">
        <div class="col-md-6 ml-auto mr-auto">
          <Card title="Create new Projection">
            <Form>
              <FormRow>
                <div class="col-8">
                  <SelectOptionInput
                    label="Check Movie"
                    v-model="MovieId"
                    :options="Movies"
                    errorMessage="Please select valid option"
                  />
                </div>
              </FormRow>
              <FormRow>
                <div class="col-8">
                  <SelectOptionInput
                    label="Check Cinema hall"
                    v-model="CinemaHallId"
                    :options="Halls"
                    errorMessage="Please select valid option"
                  />
                </div>
              </FormRow>
              <FormRow>
                <div class="col-8">
                  <DateTimePicker 
                  label="Choose Date" 
                  v-model="Slot" />
                </div>
              </FormRow>
              <FormRow>
                <div class="col-8">
                  <NumberInput 
                  label="Input price" 
                  v-model="Price" 
                  :isValid="Price >= 60 && Price <=1000"/>
                </div>
              </FormRow>
              <Button @click="handleCreate">Create</Button>
            </Form>
          </Card>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Form from "../../../components/Form/Form.vue";
import NumberInput from "../../../components/Form/NumberInput.vue"
import FormRow from "../../../components/Form/FormRow.vue";
import Button from "../../../components/Form/Button.vue";
import SelectOptionInput from "../../../components/Form/SelectOptionInput.vue";
import DateTimePicker from "../../../components/Form/DateTimePicker.vue";
import toastr from "toastr";
import { mapActions, mapGetters } from "vuex";
import Card from "../../../components/Card/Card.vue";
export default {
  components: {
    Form,
    FormRow,
    Button,
    SelectOptionInput,
    DateTimePicker,
    NumberInput,
    Card,
  },

  data: function () {
    return {
      MovieId: "",
      Slot: "",
      CinemaHallId: "",
      Price: 0,
      Movies: [],
      Halls: [],
    };
  },

  computed: {
    ...mapGetters({
      movieList: "movies/getMovies",
      hallList: "cinemaHalls/getHalls",
      result: "projections/getResult",
    }),
  },

  watch: {
    movieList(movies) {
      this.Movies = movies.map((m) => {
        return { value: m.Id, label: m.Name };
      });
    },

    hallList(halls) {
      this.Halls = halls.map((h) => {
        return { value: h.Id, label: h.Name };
      });
    },

    result({ ok, message, label }) {
      if (label !== "create") return;
      if (ok) {
        this.$router.push("/projections");
      } else {
        toastr.error(message);
      }
    },
  },

  methods: {
    ...mapActions({
      create: "projections/createProjection",
      fetchMovies: "movies/fetchMovies",
      fetchHalls: "cinemaHalls/fetchHalls",
    }),

    handleCreate() {
      const projection = {
        MovieId: this.MovieId,
        Slot: this.Slot,
        HallId: this.CinemaHallId,
        Price: this.Price
      };

      this.create(projection);
    },
  },
  mounted() {
    this.fetchMovies();
    this.fetchHalls();
  },
};
</script>
