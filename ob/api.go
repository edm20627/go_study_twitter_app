package ob

import (
	"github.com/labstack/echo"
	"github.com/objectbox/objectbox-go/objectbox"
	"go_sample/model"
	"net/http"
	"time"
)

// データ追加リクエストのパース用構造体
type AddFavoriteRequest struct {
	Name        string `josn:"name"`
	Description string `json:"description`
}

func AddFavorite(c echo.Context) error {
	favorite := new(AddFavoriteRequest)

	// リクエストパラメーターを構造体に割り当てる
	if error := c.Bind(favorite); error != nil {
		return error
	}

	box := getBox()

	const dateLayout = "2006/01/02 15:04:05"
	id, error := box.Put(&model.Favorite{
		Name:        favorite.Name,
		Description: favorite.Description,
		CreatedAt:   time.Now().Format(dateLayout),
	})
	if error != nil {
		return error
	}

	result, error := box.Get(id)
	if error != nil {
		return error
	}

	return c.JSON(http.StatusOK, result)

}

func Find(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	keyword := c.FormValue("keyword")

	if name == "" && description == "" && keyword == "" {
		return c.JSON(http.StatusNotFound, "param is not found")
	}

	box := getBox()

	if name != "" {
		result, error := findName(box, name)
		if error != nil {
			return error
		}
		return c.JSON(http.StatusOK, result)
	}

	if description != "" {
		result, error := findDescription(box, description)
		if error != nil {
			return error
		}
		return c.JSON(http.StatusOK, result)
	}

	result, error := findKeyword(box, keyword)
	if error != nil {
		return error
	}
	return c.JSON(http.StatusOK, result)
}

func findName(box *model.FavoriteBox, name string) ([]*model.Favorite, error) {
	return box.Query(model.Favorite_.Name.Contains(name, true)).Find()
}

func findDescription(box *model.FavoriteBox, description string) ([]*model.Favorite, error) {
	return box.Query(model.Favorite_.Description.Contains(description, true)).Find()
}

func findKeyword(box *model.FavoriteBox, keyword string) ([]*model.Favorite, error) {
	return box.Query(objectbox.Any(model.Favorite_.Name.Contains(keyword, true),
		model.Favorite_.Description.Contains(keyword, true))).Find()
}

func Update(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")

	if name == "" && description == "" {
		return c.JSON(http.StatusNotFound, "param is not found")
	}

	box := getBox()

	list, error := findName(box, name)
	if error != nil {
		return c.JSON(http.StatusNotFound, "item not found")
	}

	for _, item := range list {
		item.Description = description
		box.Put(item)
	}

	list, error = findName(box, name)
	if error != nil {
		return error
	}
	return c.JSON(http.StatusOK, list)
}

func GetAll(c echo.Context) error {
	box := getBox()

	var list []*model.Favorite
	list, error := box.GetAll()
	if error != nil {
		return error
	}

	return c.JSON(http.StatusOK, list)
}

func Remove(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		return c.JSON(http.StatusNotFound, "param is not found")
	}
	box := getBox()

	list, error := findName(box, name)
	if error != nil {
		return c.JSON(http.StatusNotFound, "item not found")
	}

	for _, item := range list {
		box.Remove(item)
	}

	list, error = box.GetAll()
	if error != nil {
		return error
	}
	return c.JSON(http.StatusOK, list)
}

//　テーブル呼び出し
func getBox() *model.FavoriteBox {
	return model.BoxForFavorite(InitObjectBox())
}
