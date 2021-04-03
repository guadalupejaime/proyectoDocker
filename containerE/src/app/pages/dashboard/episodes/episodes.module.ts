import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { EpisodesComponent } from '@pages/dashboard/episodes/episodes.component';
import { PageTemplateModule } from '@shared/components/page-template/page-template.module';
import { InfoEpisodeComponent } from './modals/info-episode/info-episode.component';
import { AddEpisodeComponent } from './modals/add-episode/add-episode.component';
import { MyMaterialModule } from '@shared/my-material/my-material.module';
import { ModalTemplateModule } from '@shared/components/modal-template/modal-template.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';



@NgModule({
  declarations: [
    EpisodesComponent,
    InfoEpisodeComponent,
    AddEpisodeComponent
  ],
  entryComponents: [
    InfoEpisodeComponent,
    AddEpisodeComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    PageTemplateModule,
    MyMaterialModule,
    ModalTemplateModule
  ]
})
export class EpisodesModule { }
