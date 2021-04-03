import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Character } from '@pages/dashboard/characters/characters.interface';
import { HttpService } from '@services/http.service';

@Component({
  selector: 'app-info-character',
  templateUrl: './info-character.component.html',
  styleUrls: ['./info-character.component.scss']
})
export class InfoCharacterComponent implements OnInit {
  status = 'load';
  character: Character;
  episodes = [];

  constructor(
    public sMatDialogRef: MatDialogRef<InfoCharacterComponent>,
    @Inject(MAT_DIALOG_DATA) public data: { idCharacter: string },
    private sHttpService: HttpService
  ) {
    this.getCharacter();
  }

  ngOnInit(): void {
  }

  getCharacter(): void {
    this.sHttpService.getCharacter(this.data.idCharacter).subscribe(
      (character) => {
        this.character = character;
        this.character.episode.forEach(epi => {
          let id = epi;
          if (epi.length > 2) {
            const url = epi.split('/');
            id = url[url.length - 1];
          }
          this.getNameEpisode(id);
        });
        this.status = 'ok';
      },
      (err) => {

      }
    );
  }

  getNameEpisode(id: string): void {
    this.sHttpService.getEpisode(id).subscribe(
      (episode) => {
        this.episodes.push(episode.name);
      },
      (err) => {
      }
    );
  }

}
