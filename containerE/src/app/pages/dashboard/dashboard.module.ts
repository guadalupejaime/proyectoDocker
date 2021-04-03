import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DashboardComponent } from './dashboard.component';
import { MyMaterialModule } from '@shared/my-material/my-material.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from 'src/app/app.routing';
import { CharactersModule } from '@pages/dashboard/characters/characters.module';
import { EpisodesModule } from '@pages/dashboard/episodes/episodes.module';
import { LocationsModule } from '@pages/dashboard/locations/locations.module';
import { HomeModule } from '@pages/dashboard/home/home.module';



@NgModule({
  declarations: [
    DashboardComponent
  ],
  imports: [
    CommonModule,
    MyMaterialModule,
    FormsModule,
    ReactiveFormsModule,
    AppRoutingModule,
    // Pages
    HomeModule,
    CharactersModule,
    EpisodesModule,
    LocationsModule
  ]
})
export class DashboardModule { }
