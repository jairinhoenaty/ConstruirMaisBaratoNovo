export type UserType = 'professional' | 'store';

export interface Plan {
  id: number;
  user_type: UserType;
  name: string;
  price: number;
  description: string;
  features: string; // JSON string array
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface PlanFeatures {
  features: string[];
}

// Helper function to parse features from JSON string
export const parsePlanFeatures = (featuresJson: string): string[] => {
  try {
    return JSON.parse(featuresJson);
  } catch (error) {
    console.error('Error parsing plan features:', error);
    return [];
  }
};
