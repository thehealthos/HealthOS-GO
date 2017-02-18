# Getting started

## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](https://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](https://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=HealthOS-GoLang&projectName=healthos_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=healthos_lib)

## How to Use

The following section explains how to use the HealthOS library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=healthos_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=healthos_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=healthos_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=healthos_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=healthos_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=HealthOS-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=healthos_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=healthos_lib)

## Initialization

### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| host | TODO: add a description |
| oAuthAccessToken | The OAuth 2.0 access token to be set before API calls |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


## Class Reference

### <a name="list_of_controllers"></a>List of Controllers

* [medicines_pkg](#medicines_pkg)
* [labtests_pkg](#labtests_pkg)
* [generics_pkg](#generics_pkg)
* [food_pkg](#food_pkg)
* [exercises_pkg](#exercises_pkg)
* [druginteractions_pkg](#druginteractions_pkg)
* [diseases_pkg](#diseases_pkg)
* [chat_pkg](#chat_pkg)
* [autocomplete_pkg](#autocomplete_pkg)
* [authentication_pkg](#authentication_pkg)

### <a name="medicines_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".medicines_pkg") medicines_pkg

#### Get instance

Factory for the ``` MEDICINES ``` interface can be accessed from the package medicines_pkg.

```go
medicines := medicines_pkg.NewMEDICINES()
```

#### <a name="get_manufacturer"></a>![Method: ](https://apidocs.io/img/method.png ".medicines_pkg.GetManufacturer") GetManufacturer

> Get a manufacturer by its id


```go
func (me *MEDICINES_IMPL) GetManufacturer(manufacturerId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| manufacturerId |  ``` Required ```  | Id of the manufacturer |


#### Example Usage

```go
manufacturerId := "manufacturer_id"

var result interface{}
result,_ = medicines.GetManufacturer(manufacturerId)

```


#### <a name="search_manufacturers"></a>![Method: ](https://apidocs.io/img/method.png ".medicines_pkg.SearchManufacturers") SearchManufacturers

> Search a manufacturer by its name


```go
func (me *MEDICINES_IMPL) SearchManufacturers(manufacturerQuery string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| manufacturerQuery |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
manufacturerQuery := "manufacturer_query"

var result interface{}
result,_ = medicines.SearchManufacturers(manufacturerQuery)

```


#### <a name="get_common_side_effects"></a>![Method: ](https://apidocs.io/img/method.png ".medicines_pkg.GetCommonSideEffects") GetCommonSideEffects

> Get common side effects of a medicine


```go
func (me *MEDICINES_IMPL) GetCommonSideEffects(medicineId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| medicineId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
medicineId := "medicine_id"

var result interface{}
result,_ = medicines.GetCommonSideEffects(medicineId)

```


#### <a name="get_popular_usage"></a>![Method: ](https://apidocs.io/img/method.png ".medicines_pkg.GetPopularUsage") GetPopularUsage

> Get popular usages of a medicine


```go
func (me *MEDICINES_IMPL) GetPopularUsage(medicineId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| medicineId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
medicineId := "medicine_id"

var result interface{}
result,_ = medicines.GetPopularUsage(medicineId)

```


#### <a name="get_medicines_by_manufacturer"></a>![Method: ](https://apidocs.io/img/method.png ".medicines_pkg.GetMedicinesByManufacturer") GetMedicinesByManufacturer

> Get medicines by a manufacturer


```go
func (me *MEDICINES_IMPL) GetMedicinesByManufacturer(
            page int64,
            size int64,
            manufacturerId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Required ```  | TODO: Add a parameter description |
| size |  ``` Required ```  | TODO: Add a parameter description |
| manufacturerId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("21", 10, 8)
size,_ := strconv.ParseInt("21", 10, 8)
manufacturerId := "manufacturer_id"

var result interface{}
result,_ = medicines.GetMedicinesByManufacturer(page, size, manufacturerId)

```


#### <a name="get_similar_medicines"></a>![Method: ](https://apidocs.io/img/method.png ".medicines_pkg.GetSimilarMedicines") GetSimilarMedicines

> Get similar medicines


```go
func (me *MEDICINES_IMPL) GetSimilarMedicines(medicineId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| medicineId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
medicineId := "medicine_id"

var result interface{}
result,_ = medicines.GetSimilarMedicines(medicineId)

```


#### <a name="get_medicine"></a>![Method: ](https://apidocs.io/img/method.png ".medicines_pkg.GetMedicine") GetMedicine

> Get a medicine


```go
func (me *MEDICINES_IMPL) GetMedicine(medicineId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| medicineId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
medicineId := "medicine_id"

var result interface{}
result,_ = medicines.GetMedicine(medicineId)

```


#### <a name="get_all_medicines"></a>![Method: ](https://apidocs.io/img/method.png ".medicines_pkg.GetAllMedicines") GetAllMedicines

> Get all medicines


```go
func (me *MEDICINES_IMPL) GetAllMedicines(
            page int64,
            size int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Required ```  | TODO: Add a parameter description |
| size |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("1", 10, 8)
size,_ := strconv.ParseInt("10", 10, 8)

var result interface{}
result,_ = medicines.GetAllMedicines(page, size)

```


#### <a name="search_medicines"></a>![Method: ](https://apidocs.io/img/method.png ".medicines_pkg.SearchMedicines") SearchMedicines

> Search a medicine by name


```go
func (me *MEDICINES_IMPL) SearchMedicines(medicineQuery string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| medicineQuery |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
medicineQuery := "medicine_query"

var result interface{}
result,_ = medicines.SearchMedicines(medicineQuery)

```


#### <a name="get_all_manufacturers"></a>![Method: ](https://apidocs.io/img/method.png ".medicines_pkg.GetAllManufacturers") GetAllManufacturers

> Get all manufacturers


```go
func (me *MEDICINES_IMPL) GetAllManufacturers(
            page int64,
            size int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Required ```  | TODO: Add a parameter description |
| size |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("1", 10, 8)
size,_ := strconv.ParseInt("10", 10, 8)

var result interface{}
result,_ = medicines.GetAllManufacturers(page, size)

```


#### <a name="get_alternative_medicines"></a>![Method: ](https://apidocs.io/img/method.png ".medicines_pkg.GetAlternativeMedicines") GetAlternativeMedicines

> Get substitutes of a medicine


```go
func (me *MEDICINES_IMPL) GetAlternativeMedicines(
            page int64,
            size int64,
            medicineId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Required ```  | TODO: Add a parameter description |
| size |  ``` Required ```  | TODO: Add a parameter description |
| medicineId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("21", 10, 8)
size,_ := strconv.ParseInt("21", 10, 8)
medicineId := "medicine_id"

var result interface{}
result,_ = medicines.GetAlternativeMedicines(page, size, medicineId)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="labtests_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".labtests_pkg") labtests_pkg

#### Get instance

Factory for the ``` LABTESTS ``` interface can be accessed from the package labtests_pkg.

```go
labTests := labtests_pkg.NewLABTESTS()
```

#### <a name="get_lab_test"></a>![Method: ](https://apidocs.io/img/method.png ".labtests_pkg.GetLabTest") GetLabTest

> Get labtest by id


```go
func (me *LABTESTS_IMPL) GetLabTest(labTestId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| labTestId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
labTestId := "lab_test_id"

var result interface{}
result,_ = labTests.GetLabTest(labTestId)

```


#### <a name="get_all_lab_tests"></a>![Method: ](https://apidocs.io/img/method.png ".labtests_pkg.GetAllLabTests") GetAllLabTests

> All Lab tests


```go
func (me *LABTESTS_IMPL) GetAllLabTests(
            page int64,
            size int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Required ```  | TODO: Add a parameter description |
| size |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("21", 10, 8)
size,_ := strconv.ParseInt("21", 10, 8)

var result interface{}
result,_ = labTests.GetAllLabTests(page, size)

```


#### <a name="search_lab_tests"></a>![Method: ](https://apidocs.io/img/method.png ".labtests_pkg.SearchLabTests") SearchLabTests

> Search a lab test by name


```go
func (me *LABTESTS_IMPL) SearchLabTests(labTestQuery string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| labTestQuery |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
labTestQuery := "lab_test_query"

var result interface{}
result,_ = labTests.SearchLabTests(labTestQuery)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="generics_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".generics_pkg") generics_pkg

#### Get instance

Factory for the ``` GENERICS ``` interface can be accessed from the package generics_pkg.

```go
generics := generics_pkg.NewGENERICS()
```

#### <a name="get_generic"></a>![Method: ](https://apidocs.io/img/method.png ".generics_pkg.GetGeneric") GetGeneric

> Get a generic by id


```go
func (me *GENERICS_IMPL) GetGeneric(genericId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| genericId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
genericId := "generic_id"

var result interface{}
result,_ = generics.GetGeneric(genericId)

```


#### <a name="get_medicines_by_generic"></a>![Method: ](https://apidocs.io/img/method.png ".generics_pkg.GetMedicinesByGeneric") GetMedicinesByGeneric

> Get medicines containing the generic


```go
func (me *GENERICS_IMPL) GetMedicinesByGeneric(
            page int64,
            size int64,
            exclusive bool,
            genericId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Required ```  | TODO: Add a parameter description |
| size |  ``` Required ```  | TODO: Add a parameter description |
| exclusive |  ``` Required ```  | TODO: Add a parameter description |
| genericId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("21", 10, 8)
size,_ := strconv.ParseInt("21", 10, 8)
exclusive := false
genericId := "generic_id"

var result interface{}
result,_ = generics.GetMedicinesByGeneric(page, size, exclusive, genericId)

```


#### <a name="get_all_generics"></a>![Method: ](https://apidocs.io/img/method.png ".generics_pkg.GetAllGenerics") GetAllGenerics

> All generics


```go
func (me *GENERICS_IMPL) GetAllGenerics(
            page int64,
            size int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Required ```  | TODO: Add a parameter description |
| size |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("21", 10, 8)
size,_ := strconv.ParseInt("21", 10, 8)

var result interface{}
result,_ = generics.GetAllGenerics(page, size)

```


#### <a name="search_generics"></a>![Method: ](https://apidocs.io/img/method.png ".generics_pkg.SearchGenerics") SearchGenerics

> Search a generic by name


```go
func (me *GENERICS_IMPL) SearchGenerics(genericQuery string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| genericQuery |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
genericQuery := "generic_query"

var result interface{}
result,_ = generics.SearchGenerics(genericQuery)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="food_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".food_pkg") food_pkg

#### Get instance

Factory for the ``` FOOD ``` interface can be accessed from the package food_pkg.

```go
food := food_pkg.NewFOOD()
```

#### <a name="search_food_restaurants"></a>![Method: ](https://apidocs.io/img/method.png ".food_pkg.SearchFoodRestaurants") SearchFoodRestaurants

> TODO: Add Description


```go
func (me *FOOD_IMPL) SearchFoodRestaurants(foodRestaurantQuery string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| foodRestaurantQuery |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
foodRestaurantQuery := "food_restaurant_query"

var result interface{}
result,_ = food.SearchFoodRestaurants(foodRestaurantQuery)

```


#### <a name="get_all_food_brands"></a>![Method: ](https://apidocs.io/img/method.png ".food_pkg.GetAllFoodBrands") GetAllFoodBrands

> TODO: Add Description


```go
func (me *FOOD_IMPL) GetAllFoodBrands(
            page int64,
            size int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Required ```  | TODO: Add a parameter description |
| size |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("1", 10, 8)
size,_ := strconv.ParseInt("10", 10, 8)

var result interface{}
result,_ = food.GetAllFoodBrands(page, size)

```


#### <a name="get_food_items_by_restaurant"></a>![Method: ](https://apidocs.io/img/method.png ".food_pkg.GetFoodItemsByRestaurant") GetFoodItemsByRestaurant

> TODO: Add Description


```go
func (me *FOOD_IMPL) GetFoodItemsByRestaurant(
            page int64,
            size int64,
            foodRestaurantId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Required ```  | TODO: Add a parameter description |
| size |  ``` Required ```  | TODO: Add a parameter description |
| foodRestaurantId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("21", 10, 8)
size,_ := strconv.ParseInt("21", 10, 8)
foodRestaurantId := "food_restaurant_id"

var result interface{}
result,_ = food.GetFoodItemsByRestaurant(page, size, foodRestaurantId)

```


#### <a name="search_food_brands"></a>![Method: ](https://apidocs.io/img/method.png ".food_pkg.SearchFoodBrands") SearchFoodBrands

> TODO: Add Description


```go
func (me *FOOD_IMPL) SearchFoodBrands(foodBrandQuery string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| foodBrandQuery |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
foodBrandQuery := "food_brand_query"

var result interface{}
result,_ = food.SearchFoodBrands(foodBrandQuery)

```


#### <a name="get_food_item"></a>![Method: ](https://apidocs.io/img/method.png ".food_pkg.GetFoodItem") GetFoodItem

> TODO: Add Description


```go
func (me *FOOD_IMPL) GetFoodItem(foodItemId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| foodItemId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
foodItemId := "food_item_id"

var result interface{}
result,_ = food.GetFoodItem(foodItemId)

```


#### <a name="get_all_food_items"></a>![Method: ](https://apidocs.io/img/method.png ".food_pkg.GetAllFoodItems") GetAllFoodItems

> TODO: Add Description


```go
func (me *FOOD_IMPL) GetAllFoodItems(
            page int64,
            size int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Required ```  | TODO: Add a parameter description |
| size |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("21", 10, 8)
size,_ := strconv.ParseInt("21", 10, 8)

var result interface{}
result,_ = food.GetAllFoodItems(page, size)

```


#### <a name="search_food_items"></a>![Method: ](https://apidocs.io/img/method.png ".food_pkg.SearchFoodItems") SearchFoodItems

> TODO: Add Description


```go
func (me *FOOD_IMPL) SearchFoodItems(foodItemQuery string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| foodItemQuery |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
foodItemQuery := "food_item_query"

var result interface{}
result,_ = food.SearchFoodItems(foodItemQuery)

```


#### <a name="get_all_food_restaurants"></a>![Method: ](https://apidocs.io/img/method.png ".food_pkg.GetAllFoodRestaurants") GetAllFoodRestaurants

> TODO: Add Description


```go
func (me *FOOD_IMPL) GetAllFoodRestaurants(
            page int64,
            size int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Required ```  | TODO: Add a parameter description |
| size |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("1", 10, 8)
size,_ := strconv.ParseInt("10", 10, 8)

var result interface{}
result,_ = food.GetAllFoodRestaurants(page, size)

```


#### <a name="get_food_items_by_brand"></a>![Method: ](https://apidocs.io/img/method.png ".food_pkg.GetFoodItemsByBrand") GetFoodItemsByBrand

> TODO: Add Description


```go
func (me *FOOD_IMPL) GetFoodItemsByBrand(
            page int64,
            size int64,
            foodBrandId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Required ```  | TODO: Add a parameter description |
| size |  ``` Required ```  | TODO: Add a parameter description |
| foodBrandId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("21", 10, 8)
size,_ := strconv.ParseInt("21", 10, 8)
foodBrandId := "food_brand_id"

var result interface{}
result,_ = food.GetFoodItemsByBrand(page, size, foodBrandId)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="exercises_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".exercises_pkg") exercises_pkg

#### Get instance

Factory for the ``` EXERCISES ``` interface can be accessed from the package exercises_pkg.

```go
exercises := exercises_pkg.NewEXERCISES()
```

#### <a name="get_exercise"></a>![Method: ](https://apidocs.io/img/method.png ".exercises_pkg.GetExercise") GetExercise

> TODO: Add Description


```go
func (me *EXERCISES_IMPL) GetExercise(exerciseId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| exerciseId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
exerciseId := "exercise_id"

var result interface{}
result,_ = exercises.GetExercise(exerciseId)

```


#### <a name="get_all_exercises"></a>![Method: ](https://apidocs.io/img/method.png ".exercises_pkg.GetAllExercises") GetAllExercises

> TODO: Add Description


```go
func (me *EXERCISES_IMPL) GetAllExercises(
            page int64,
            size int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Required ```  | TODO: Add a parameter description |
| size |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("1", 10, 8)
size,_ := strconv.ParseInt("10", 10, 8)

var result interface{}
result,_ = exercises.GetAllExercises(page, size)

```


#### <a name="search_exercises"></a>![Method: ](https://apidocs.io/img/method.png ".exercises_pkg.SearchExercises") SearchExercises

> TODO: Add Description


```go
func (me *EXERCISES_IMPL) SearchExercises(exerciseQuery string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| exerciseQuery |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
exerciseQuery := "exercise_query"

var result interface{}
result,_ = exercises.SearchExercises(exerciseQuery)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="druginteractions_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".druginteractions_pkg") druginteractions_pkg

#### Get instance

Factory for the ``` DRUGINTERACTIONS ``` interface can be accessed from the package druginteractions_pkg.

```go
drugInteractions := druginteractions_pkg.NewDRUGINTERACTIONS()
```

#### <a name="get_generics_interactions"></a>![Method: ](https://apidocs.io/img/method.png ".druginteractions_pkg.GetGenericsInteractions") GetGenericsInteractions

> TODO: Add Description


```go
func (me *DRUGINTERACTIONS_IMPL) GetGenericsInteractions(genericId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| genericId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
genericId := "generic_id"

var result interface{}
result,_ = drugInteractions.GetGenericsInteractions(genericId)

```


#### <a name="get_medicine_interactions"></a>![Method: ](https://apidocs.io/img/method.png ".druginteractions_pkg.GetMedicineInteractions") GetMedicineInteractions

> TODO: Add Description


```go
func (me *DRUGINTERACTIONS_IMPL) GetMedicineInteractions(medicineId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| medicineId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
medicineId := "medicine_id"

var result interface{}
result,_ = drugInteractions.GetMedicineInteractions(medicineId)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="diseases_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".diseases_pkg") diseases_pkg

#### Get instance

Factory for the ``` DISEASES ``` interface can be accessed from the package diseases_pkg.

```go
diseases := diseases_pkg.NewDISEASES()
```

#### <a name="get_disease"></a>![Method: ](https://apidocs.io/img/method.png ".diseases_pkg.GetDisease") GetDisease

> TODO: Add Description


```go
func (me *DISEASES_IMPL) GetDisease(diseaseId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| diseaseId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
diseaseId := "disease_id"

var result interface{}
result,_ = diseases.GetDisease(diseaseId)

```


#### <a name="get_all_diseases"></a>![Method: ](https://apidocs.io/img/method.png ".diseases_pkg.GetAllDiseases") GetAllDiseases

> TODO: Add Description


```go
func (me *DISEASES_IMPL) GetAllDiseases(
            page int64,
            size int64)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Required ```  | TODO: Add a parameter description |
| size |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
page,_ := strconv.ParseInt("1", 10, 8)
size,_ := strconv.ParseInt("10", 10, 8)

var result interface{}
result,_ = diseases.GetAllDiseases(page, size)

```


#### <a name="search_diseases"></a>![Method: ](https://apidocs.io/img/method.png ".diseases_pkg.SearchDiseases") SearchDiseases

> TODO: Add Description


```go
func (me *DISEASES_IMPL) SearchDiseases(diseaseQuery string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| diseaseQuery |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
diseaseQuery := "disease_query"

var result interface{}
result,_ = diseases.SearchDiseases(diseaseQuery)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="chat_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".chat_pkg") chat_pkg

#### Get instance

Factory for the ``` CHAT ``` interface can be accessed from the package chat_pkg.

```go
chat := chat_pkg.NewCHAT()
```

#### <a name="get_excercises_chat"></a>![Method: ](https://apidocs.io/img/method.png ".chat_pkg.GetExcercisesChat") GetExcercisesChat

> TODO: Add Description


```go
func (me *CHAT_IMPL) GetExcercisesChat()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = chat.GetExcercisesChat()

```


#### <a name="get_food_items_chat"></a>![Method: ](https://apidocs.io/img/method.png ".chat_pkg.GetFoodItemsChat") GetFoodItemsChat

> TODO: Add Description


```go
func (me *CHAT_IMPL) GetFoodItemsChat()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = chat.GetFoodItemsChat()

```


#### <a name="get_medicine_chat"></a>![Method: ](https://apidocs.io/img/method.png ".chat_pkg.GetMedicineChat") GetMedicineChat

> TODO: Add Description


```go
func (me *CHAT_IMPL) GetMedicineChat()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = chat.GetMedicineChat()

```


[Back to List of Controllers](#list_of_controllers)

### <a name="autocomplete_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".autocomplete_pkg") autocomplete_pkg

#### Get instance

Factory for the ``` AUTOCOMPLETE ``` interface can be accessed from the package autocomplete_pkg.

```go
autocomplete := autocomplete_pkg.NewAUTOCOMPLETE()
```

#### <a name="get_exercises"></a>![Method: ](https://apidocs.io/img/method.png ".autocomplete_pkg.GetExercises") GetExercises

> TODO: Add Description


```go
func (me *AUTOCOMPLETE_IMPL) GetExercises(exerciseQuery string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| exerciseQuery |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
exerciseQuery := "exercise_query"

var result interface{}
result,_ = autocomplete.GetExercises(exerciseQuery)

```


#### <a name="get_diseases"></a>![Method: ](https://apidocs.io/img/method.png ".autocomplete_pkg.GetDiseases") GetDiseases

> TODO: Add Description


```go
func (me *AUTOCOMPLETE_IMPL) GetDiseases(diseaseQuery string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| diseaseQuery |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
diseaseQuery := "disease_query"

var result interface{}
result,_ = autocomplete.GetDiseases(diseaseQuery)

```


#### <a name="get_lab_tests"></a>![Method: ](https://apidocs.io/img/method.png ".autocomplete_pkg.GetLabTests") GetLabTests

> TODO: Add Description


```go
func (me *AUTOCOMPLETE_IMPL) GetLabTests(labTestQuery string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| labTestQuery |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
labTestQuery := "lab_test_query"

var result interface{}
result,_ = autocomplete.GetLabTests(labTestQuery)

```


#### <a name="get_food_items"></a>![Method: ](https://apidocs.io/img/method.png ".autocomplete_pkg.GetFoodItems") GetFoodItems

> TODO: Add Description


```go
func (me *AUTOCOMPLETE_IMPL) GetFoodItems(foodItemQuery string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| foodItemQuery |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
foodItemQuery := "food_item_query"

var result interface{}
result,_ = autocomplete.GetFoodItems(foodItemQuery)

```


#### <a name="get_generics"></a>![Method: ](https://apidocs.io/img/method.png ".autocomplete_pkg.GetGenerics") GetGenerics

> TODO: Add Description


```go
func (me *AUTOCOMPLETE_IMPL) GetGenerics(genericQuery string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| genericQuery |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
genericQuery := "generic_query"

var result interface{}
result,_ = autocomplete.GetGenerics(genericQuery)

```


#### <a name="get_medicines"></a>![Method: ](https://apidocs.io/img/method.png ".autocomplete_pkg.GetMedicines") GetMedicines

> TODO: Add Description


```go
func (me *AUTOCOMPLETE_IMPL) GetMedicines(medicineQuery string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| medicineQuery |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
medicineQuery := "medicine_query"

var result interface{}
result,_ = autocomplete.GetMedicines(medicineQuery)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="authentication_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".authentication_pkg") authentication_pkg

#### Get instance

Factory for the ``` AUTHENTICATION ``` interface can be accessed from the package authentication_pkg.

```go
authentication := authentication_pkg.NewAUTHENTICATION()
```

#### <a name="create_request_access_token"></a>![Method: ](https://apidocs.io/img/method.png ".authentication_pkg.CreateRequestAccessToken") CreateRequestAccessToken

> *Tags:*  ``` Skips Authentication ``` 

> TODO: Add Description


```go
func (me *AUTHENTICATION_IMPL) CreateRequestAccessToken(body *models_pkg.RequestAccessTokenRequest)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
bodyValue := []byte("{\n  \"grant_type\": \"client_credentials\",\n  \"client_id\": \"{{client_id}}\",\n  \"client_secret\": \"{{client_secret}}\",\n  \"scope\": \"public read write\"\n}")
var body *models_pkg.RequestAccessTokenRequest
json.Unmarshal(bodyValue,&body)

var result interface{}
result,_ = authentication.CreateRequestAccessToken(body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 401 | Unauthorized |



[Back to List of Controllers](#list_of_controllers)



