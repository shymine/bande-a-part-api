"""
URL configuration for bande_a_part_api project.

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/5.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include
from django.views.generic import TemplateView

from api.views import LibraryDetail, LibraryListView, BookListDetail, BookListListView, AuthorDetail, AuthorListView, BookDetail, BookListView, EditorDetail, EditorListView, GenreDetail, GenreListView

urlpatterns = [
    path('admin/', admin.site.urls),

    path('author/', AuthorListView.as_view()),
    path('author/<int:pk>/', AuthorDetail.as_view()),

    path('book/', BookListView.as_view()),
    path('book/<int:pk>/', BookDetail.as_view()),

    path('editor/', EditorListView.as_view()),
    path('editor/<int:pk>/', EditorDetail.as_view()),

    path('genre/', GenreListView.as_view()),
    path('genre/<int:pk>/', GenreDetail.as_view()),

    path('bookList/', BookListListView.as_view()),
    path('bookList/<int:pk>/', BookListDetail.as_view()),

    path('library/', LibraryListView.as_view()),
    path('library/<int:pk>/', LibraryDetail.as_view()),

    path('home/', TemplateView.as_view(template_name="home.html"), name="home"),
    path('api-auth/', include('rest_framework.urls', namespace='rest_framework')),
]
