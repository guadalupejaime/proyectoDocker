import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Episode } from '@pages/dashboard/episodes/episodes.interface';
import { reverseMonths } from '@pages/dashboard/characters/characters.models';
import { HttpService } from '@services/http.service';

@Component({
  selector: 'app-info-episode',
  templateUrl: './info-episode.component.html',
  styleUrls: ['./info-episode.component.scss']
})
export class InfoEpisodeComponent implements OnInit {
  status = 'load';
  episode: Episode;
  characters = [];

  constructor(
    public sMatDialogRef: MatDialogRef<InfoEpisodeComponent>,
    @Inject(MAT_DIALOG_DATA) public data: { idEpisode: string },
    private sHttpService: HttpService
  ) {
    this.getEpisode();
  }

  ngOnInit(): void {
  }

  getEpisode(): void {
    this.sHttpService.getEpisode(this.data.idEpisode).subscribe(
      (episode) => {
        this.episode = episode;
        this.episode.characters.forEach(ch => {
          let id = ch;
          if (ch.length > 2) {
            const url = ch.split('/');
            id = url[url.length - 1];
          }
          this.getNameCharacter(id);
        });
        this.status = 'ok';
      },
      (err) => {

      }
    );
  }

  getNameCharacter(id: string): void {
    this.sHttpService.getCharacter(id).subscribe(
      (character) => {
        this.characters.push(character.name);
      },
      (err) => {
      }
    );
  }

}
