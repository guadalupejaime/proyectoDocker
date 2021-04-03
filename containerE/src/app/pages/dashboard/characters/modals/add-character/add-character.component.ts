import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { Location } from '@pages/dashboard/locations/locations.interface';
import { Episode } from '@pages/dashboard/episodes/episodes.interface';
import { listGender, listSpecies, listStatus } from '@pages/dashboard/characters/characters.models';
import { HttpService } from '@services/http.service';

@Component({
  selector: 'app-add-character',
  templateUrl: './add-character.component.html',
  styleUrls: ['./add-character.component.scss']
})
export class AddCharacterComponent implements OnInit {
  @ViewChild('episodeInput') episodeInput: ElementRef<HTMLInputElement>;
  allStatus = listStatus;
  allSpecies = listSpecies;
  allGenders = listGender;
  locations: Location[];
  episodes: Episode[] = [];
  filteredOptions: Location[] = [];
  filteredEpisodes: Episode[] = [];
  ctrlLocations = new FormControl(null, Validators.required);
  ctrlType = new FormControl(null, Validators.required);
  ctrlOrigin = new FormControl(null, Validators.required);
  ctrlEpisode = new FormControl([], Validators.required);
  selectedEpisodes: Episode[] = [];
  formCharacters: FormGroup;
  checked = false;

  constructor(
    public sMatDialogRef: MatDialogRef<AddCharacterComponent>,
    private sHttpService: HttpService
  ) {
    this.formCharacters = new FormGroup({
      name: new FormControl(null, Validators.required),
      status: new FormControl(null, Validators.required),
      species: new FormControl(null, Validators.required),
      gender: new FormControl(null, Validators.required),
      image: new FormControl('https://rickandmortyapi.com/api/character/avatar/19.jpeg', Validators.required),
    });
  }

  ngOnInit(): void {
    this.getLocations();
    this.getEpisodes();

    this.ctrlLocations.valueChanges.subscribe(
      (location: string) => {
        const temp = typeof location;
        if (temp === 'string') {
          this.filterLocations(location.toLowerCase());
        }
      });

    this.ctrlOrigin.valueChanges.subscribe(
      (origin: string) => {
        const temp = typeof origin;
        if (temp === 'string') {
          this.filterLocations(origin.toLowerCase());
        }
      });

    this.ctrlEpisode.valueChanges.subscribe(
      (episode: string) => {
        const temp = typeof episode;
        if (temp === 'string') {
          this.filterEpisodes(episode.toLowerCase());
        }
      });

  }

  // get locations list
  private getLocations(): void {
    this.sHttpService.getLocations({ limit: 500, offset: 0 }).subscribe(
      (locations) => {
        this.locations = locations.results;
      },
      (err) => {

      }
    );
  }

  // get locations list
  private getEpisodes(): void {
    this.sHttpService.getEpisodes({ limit: 500, offset: 0 }).subscribe(
      (episodes) => {
        episodes.results.forEach(episode => {
          const newEpisode: Episode = { ...episode, checked: false };
          this.episodes.push(newEpisode);
        });
        // this.episodes = episodes.results;
      },
      (err) => {

      }
    );
  }

  private filterLocations(myFilter?: string): void {
    if (myFilter) {
      this.filteredOptions = this.locations.filter((location) => {
        return location.name.toLowerCase().indexOf(myFilter) === 0;
      });
    } else {
      if (this.locations.length < 15) {
        this.filteredOptions = this.locations;
      }
    }
  }

  private filterEpisodes(myFilter?: string): void {
    if (myFilter) {
      this.filteredEpisodes = this.episodes.filter((episode) => {
        return episode.name.toLowerCase().indexOf(myFilter) === 0;
      });
    } else {
      if (this.locations.length < 15) {
        this.filteredEpisodes = this.episodes;
      }
    }
  }

  displayFn(location: Location): string {
    return location && location.name ? location.name : '';
  }

  displayFnEpisode(episode: Episode): string {
    return episode && episode.name ? episode.name : '';
  }

  filterLocationSelected(location: string): void {
    this.ctrlLocations.setValue(location);
  }

  filterOriginSelected(origin: string): void {
    this.ctrlOrigin.setValue(origin);
  }

  filterEpisodeSelected(episode: Episode): void {
    const index = this.filteredEpisodes.indexOf(episode);
    if (index >= 0) {
      this.filteredEpisodes[index].checked = true;
    }
    if (this.selectedEpisodes.indexOf(episode) < 0) {
      this.selectedEpisodes.push(episode);
    }
    this.ctrlEpisode.setValue(this.selectedEpisodes);
  }


  removeEpisode(episode: Episode): void {
    const index = this.selectedEpisodes.indexOf(episode);

    if (index >= 0) {
      this.selectedEpisodes.splice(index, 1);
      this.filteredEpisodes[index].checked = false;
    }
  }

  changeEpisodes(episode: Episode): void {
    const index = this.filteredEpisodes.indexOf(episode);
    this.episodeInput.nativeElement.value = '';
    if (index >= 0) {
      if (this.filteredEpisodes[index].checked) {
        this.removeEpisode(episode);
        return;
      }
      this.filterEpisodeSelected(episode);
    }
  }

  isValidForm(): boolean {
    const validEpisode = this.ctrlEpisode.value.length > 0;
    this.formCharacters.markAllAsTouched();
    this.ctrlEpisode.markAsTouched();
    this.ctrlLocations.markAsTouched();
    this.ctrlOrigin.markAsTouched();
    this.ctrlType.markAsTouched();
    return this.formCharacters.valid && this.ctrlLocations.valid &&
      this.ctrlOrigin.valid && validEpisode && this.ctrlType.valid;
  }

  postCharacter(): void {
    if (this.isValidForm()) {
      const episodesSelected = [];
      for (const episode of this.ctrlEpisode.value) {
        episodesSelected.push(episode.id + '');
      }
      const newCharacter = {
        ...this.formCharacters.value,
        type: this.ctrlType.value,
        origin: {
          name: this.ctrlOrigin.value.name,
          url: this.ctrlOrigin.value.id + ''
        },
        location: {
          name: this.ctrlLocations.value.name,
          url: this.ctrlLocations.value.id + '',
        },
        episode: episodesSelected
      };

      this.sHttpService.postCharacters(newCharacter).subscribe(
        (resp) => {
          this.sMatDialogRef.close(true);
        },
        (err) => {
          this.sMatDialogRef.close('error');
        }
      );
    }
  }
}
