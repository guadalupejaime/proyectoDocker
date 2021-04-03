import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { EpisodesComponent } from '@pages/dashboard/episodes/episodes.component';
import { PageTemplateModule } from '@shared/components/page-template/page-template.module';



@NgModule({
  declarations: [
    EpisodesComponent
  ],
  imports: [
    CommonModule,
    PageTemplateModule
  ]
})
export class EpisodesModule { }
