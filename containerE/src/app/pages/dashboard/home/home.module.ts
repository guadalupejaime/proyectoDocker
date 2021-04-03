import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from '@pages/dashboard/home/home.component';
import { CardComponent } from '@shared/components/card/card.component';
import { MyMaterialModule } from '@shared/my-material/my-material.module';
import { AppRoutingModule } from 'src/app/app.routing';
import { PageTemplateModule } from '@shared/components/page-template/page-template.module';



@NgModule({
  declarations: [
    HomeComponent,
    CardComponent
  ],
  imports: [
    CommonModule,
    MyMaterialModule,
    AppRoutingModule,
    PageTemplateModule
  ]
})
export class HomeModule { }
