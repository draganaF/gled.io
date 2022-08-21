package service

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	apicontract "github.com/draganaF/gled.io/ProjectionService/apiContract"
	"github.com/draganaF/gled.io/ProjectionService/model"
	"github.com/draganaF/gled.io/ProjectionService/repository"
	"github.com/draganaF/gled.io/ProjectionService/utils"
)

type ProjectionService struct {
	repository *repository.ProjectionRepository
}

func NewProjectionService() *ProjectionService {
	return &ProjectionService{
		repository: &repository.ProjectionRepository{
			DB: utils.DB,
		},
	}
}

func (projectionService *ProjectionService) Search(params *apicontract.SearchParams) (*[]model.Projection, error) {
	projections := projectionService.repository.Search(params)

	if projections == nil {
		return nil, errors.New("no projections found")
	}

	if params.Score != 0 {
		newProjections := projectionService.GetScore(params.Score, projections)

		return newProjections, nil

	}
	return projections, nil
}

func (projectionService *ProjectionService) GetScore(score int, projections *[]model.Projection) *[]model.Projection {
	newProjections := []model.Projection{}

	for _, projection := range *projections {
		pScore := projectionService.RemoteCall(int(projection.Movie.Id))
		if score == pScore {
			newProjections = append(newProjections, projection)
		}
	}

	return &newProjections
}

func (projectionService *ProjectionService) RemoteCall(movieId int) int {
	client := http.Client{}
	url := os.Getenv("RECENSION_SERVICE_URL") + "/api/recensions/score-by-movie-id/" + strconv.Itoa(movieId)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")

	res, _ := client.Do(req)
	defer res.Body.Close()

	bodyBytes, _ := io.ReadAll(res.Body)

	score, _ := strconv.Atoi(string(string(bodyBytes)[15]))
	return score
}

func (projectionService *ProjectionService) ReadAll() (*[]model.Projection, error) {

	projections := projectionService.repository.ReadAll()

	if projections == nil {
		return nil, errors.New("No projections found")
	}

	return projections, nil
}

func (projectionService *ProjectionService) ReadProjectionById(id uint) (*model.Projection, error) {

	projection := projectionService.repository.ReadProjectionById(id)

	if projection == nil {
		return nil, errors.New("There is no projection with ID " + strconv.FormatUint(uint64(id), 10))
	}

	return projection, nil
}

func (projectionService *ProjectionService) ReadProjectionsByMovieId(id uint) (*[]model.Projection, error) {

	projection := projectionService.repository.ReadProjectionsByMovieId(id)

	if projection == nil {
		return nil, errors.New("There is no user with ID " + strconv.FormatUint(uint64(id), 10))
	}

	return projection, nil
}

func (projectionService *ProjectionService) Create(projectionDto apicontract.CreateProjectionRequest) (*model.Projection, error) {
	projection := projectionDto.ToProjection()

	movieService := NewMovieService()
	movie := movieService.repository.ReadMovieById(projection.MovieId)
	if movie == nil {
		return nil, errors.New("movie does not exist")
	}
	projection.Movie = *movie

	cinemaHallService := NewCinemaHallService()
	cinemaHall := cinemaHallService.repository.ReadCinemaHallById(projection.CinemaHallId)

	if cinemaHall == nil {
		return nil, errors.New("cinema hall does not exist")
	}
	projection.CinemaHall = *cinemaHall

	projections := projectionService.repository.ReadAll()

	if projections != nil {
		for _, value := range *projections {
			endDate1 := projection.Slot.Add(time.Minute * time.Duration(projection.Movie.Duration))
			endDate2 := value.Slot.Add(time.Minute * time.Duration(value.Movie.Duration))

			if projection.CinemaHallId == value.CinemaHallId {
				if projectionService.AreTwoDateRangeOverlapping(projection.Slot, endDate1, value.Slot, endDate2) {
					return nil, errors.New("time slot in that hall is not available")
				}
			}
		}
	}

	projection = projectionService.repository.Create(projection)

	if projection == nil {
		return nil, errors.New("something went wrong")
	}

	return projection, nil
}

func (projectionService *ProjectionService) Update(projectionDto apicontract.UpdateProjectionRequest) (*model.Projection, error) {
	existingProjection := projectionService.repository.ReadProjectionById(projectionDto.Id)

	if existingProjection == nil {
		return nil, errors.New("No projection found with ID " + strconv.FormatUint(uint64(projectionDto.Id), 10))
	}

	existingProjection.Movie = projectionDto.Movie
	existingProjection.Slot = projectionDto.Slot
	existingProjection.CinemaHall = projectionDto.CinemaHall

	savedProjection := projectionService.repository.Update(existingProjection)

	if savedProjection == nil {
		return nil, errors.New("something went wrong")
	}

	return savedProjection, nil
}

func (projectionService *ProjectionService) Delete(id uint) error {
	_, err := projectionService.ReadProjectionById(id)
	if err != nil {
		return err
	}
	projectionService.repository.Delete(id)
	return nil
}

func (projectionService *ProjectionService) AreTwoDateRangeOverlapping(startDate1 time.Time, endDate1 time.Time, startDate2 time.Time, endDate2 time.Time) bool {
	if AfterOrEquals(startDate1, startDate2) && BeforeOrEquals(endDate1, endDate2) {
		return true
	}

	if BeforeOrEquals(startDate1, startDate2) && AfterOrEquals(endDate1, endDate2) {
		return true
	}

	if AfterOrEquals(startDate1, startDate2) && AfterOrEquals(endDate1, endDate2) && BeforeOrEquals(startDate1, endDate2) {
		return true
	}

	if AfterOrEquals(startDate2, startDate1) && AfterOrEquals(endDate2, endDate1) && BeforeOrEquals(startDate2, endDate1) {
		return true
	}

	return false
}

func AfterOrEquals(t1 time.Time, t2 time.Time) bool {
	return t1.After(t2) || t1.Equal(t2)
}

func BeforeOrEquals(t1 time.Time, t2 time.Time) bool {
	return t1.Before(t2) || t1.Equal(t2)
}
