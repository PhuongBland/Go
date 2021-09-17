// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}

func (_ tApp) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.List", args).URL
}

func (_ tApp) Add(
		text string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "text", text)
	return revel.MainRouter.Reverse("App.Add", args).URL
}

func (_ tApp) Remove(
		index int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "index", index)
	return revel.MainRouter.Reverse("App.Remove", args).URL
}

func (_ tApp) RemoveAll(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.RemoveAll", args).URL
}


type tAuthController struct {}
var AuthController tAuthController


func (_ tAuthController) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AuthController.Index", args).URL
}

func (_ tAuthController) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AuthController.Login", args).URL
}

func (_ tAuthController) LoginForm(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AuthController.LoginForm", args).URL
}


type tCategoryController struct {}
var CategoryController tCategoryController


func (_ tCategoryController) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CategoryController.Index", args).URL
}

func (_ tCategoryController) GetAll(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CategoryController.GetAll", args).URL
}

func (_ tCategoryController) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CategoryController.Create", args).URL
}


type tCourseController struct {}
var CourseController tCourseController


func (_ tCourseController) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CourseController.Index", args).URL
}

func (_ tCourseController) GetAll(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CourseController.GetAll", args).URL
}

func (_ tCourseController) InsertOne(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CourseController.InsertOne", args).URL
}

func (_ tCourseController) InsertMany(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CourseController.InsertMany", args).URL
}


type tQuestionController struct {}
var QuestionController tQuestionController


func (_ tQuestionController) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("QuestionController.Index", args).URL
}

func (_ tQuestionController) GetAll(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("QuestionController.GetAll", args).URL
}

func (_ tQuestionController) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("QuestionController.Create", args).URL
}


type tUserController struct {}
var UserController tUserController


func (_ tUserController) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.Index", args).URL
}

func (_ tUserController) GetAll(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.GetAll", args).URL
}

func (_ tUserController) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.Register", args).URL
}

func (_ tUserController) Delete(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("UserController.Delete", args).URL
}

func (_ tUserController) GetUserFormID(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("UserController.GetUserFormID", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


