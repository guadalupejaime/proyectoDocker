import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from '@pages/dashboard/home/home.component';
import { CardComponent } from '@shared/components/card/card.component';



@NgModule({
  declarations: [
    HomeComponent,
    CardComponent
  ],
  imports: [
    CommonModule
  ]
})
export class HomeModule { }
