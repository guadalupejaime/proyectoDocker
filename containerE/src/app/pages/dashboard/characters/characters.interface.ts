export interface Character {
    name: string;
    created: string;
    gender: string;
    episode: string[];
    id: number;
    image: string;
    location: Location;
    origin: Origin;
    species: string;
    status: string;
    type: string;
    checked?: boolean;
}

export interface Info {
    count: number;
    next: string;
    page: number;
    prev: string;
}

export interface Location {
    name: string;
    url: string;
}

export interface Origin {
    name: string;
    url: string;
}

export interface Characters {
    info: Info;
    results: Character[];
    total_found: number;
    total_returned: number;
}
