import { writable } from "svelte/store";

/**
 * @typedef {Object} UserData
 * @property {string} DiscordUserHandle
 * @property {string | null} SyncedGoogleID
 * @property {number} SyncedGithubID
 * @property {Date} ManualProSubExpiresAt
 * @property {Date | null} LifetimeProSubCreatedAt
 * @property {Date} MembershipExpiresAt
 * @property {string} Email
 * @property {string} Currency
 * @property {number} Xp
 * @property {number} Level
 * @property {number} XPForLevel
 * @property {number} XPTotalForLevel
 * @property {string} Role
 * @property {number} Gems
 * @property {number} Armor
 * @property {Date} CreatedAt
 * @property {Date} UpdatedAt
 * @property {string} FirstName
 * @property {string} LastName
 * @property {string} Handle
 * @property {string} Bio
 * @property {string} JobTitle
 * @property {string} Location
 * @property {string} City
 * @property {string} Country
 * @property {string} GithubHandle
 * @property {string} WebsiteURL
 * @property {string} ProfileImageURL
 * @property {boolean} IsSubscribed
 * @property {boolean} GithubSynced
 *
 */

/**
 * @typedef {Object} User
 * @property {boolean} isArchmage
 * @property {boolean} isLoggedIn
 * @property {number} seerStones
 * @property {number} bakedSalmon
 * @property {number} xpPotion
 * @property {Date|null} xpTimer
 * @property {number} frozenFlame
 * @property {UserData} userData
 */
export const User = writable({
  isArchmage: false,
  isLoggedIn: false,
  seerStones: 0,
  bakedSalmon: 0,
  xpPotion: 0,
  xpTimer: null,
  frozenFlame: 0,
  userData: {},
});
