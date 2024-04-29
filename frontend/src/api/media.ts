import { apiCall, TimeStamped } from './common';

const assetUrl = (bucket: string, asset: Asset): string => `${window.gbans.asset_url}/${bucket}/${asset.name}`;

export const assetURLMedia = (asset: Asset) => assetUrl('media', asset);

export const assetURLDemo = (asset: Asset) => assetUrl('demo', asset);

export enum MediaTypes {
    video,
    image,
    other
}

export const mediaType = (mime_type: string): MediaTypes => {
    if (mime_type.startsWith('image/')) {
        return MediaTypes.image;
    } else if (mime_type.startsWith('video/')) {
        return MediaTypes.video;
    } else {
        return MediaTypes.other;
    }
};

export interface BaseUploadedMedia extends TimeStamped {
    media_id: number;
    author_id: number;
    mime_type: string;
    size: number;
    name: string;
    contents: Uint8Array;
    deleted: boolean;
    asset: Asset;
}

export interface Asset {
    asset_id: string;
    bucket: string;
    path: string;
    name: string;
    mime_type: string;
    size: number;
    old_id: number;
}

export interface MediaUploadResponse extends BaseUploadedMedia {
    url: string;
}

export interface UserUploadedFile {
    content: string;
    name: string;
    mime: string;
    size: number;
}

export const apiSaveMedia = async (upload: UserUploadedFile) =>
    await apiCall<MediaUploadResponse, UserUploadedFile>(`/api/media`, 'POST', upload);

export const apiSaveContestEntryMedia = async (contest_id: string, upload: UserUploadedFile) =>
    await apiCall<MediaUploadResponse, UserUploadedFile>(`/api/contests/${contest_id}/upload`, 'POST', upload);
