import axios from "axios";
import { getMapsApiOptions } from '../jsm/load-maps-api';
const { apiKey, mapId } = getMapsApiOptions();
const getImageMap = (center, zoom, size, scale, mapType) => {
    mapType = mapType ? mapType : "satellite"
    return `https://maps.googleapis.com/maps/api/staticmap?center=${center.lat},${center.lng}&zoom=${zoom}&size=${size}x${size}&maptype=${mapType}&scale=${scale}&key=${apiKey}`
}
export {
    getImageMap
}